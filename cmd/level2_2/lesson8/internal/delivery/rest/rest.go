package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"gb/cmd/level2_2/lesson8/internal/service"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Rest struct {
	server *http.Server
	games  sync.Map
}

func NewRest(listen string, size int) (*Rest, error) {

	rest := &Rest{}
	router := http.NewServeMux()

	router.HandleFunc("/start", func(writer http.ResponseWriter, request *http.Request) {
		defer rest.save()
		game, err := service.NewTicTacToe(size)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}
		session := rand.Int63()
		rest.games.Store(session, game)

		writer.Write([]byte(fmt.Sprintf("%d", session)))
		writer.WriteHeader(http.StatusOK)

	})

	router.HandleFunc("/move", func(writer http.ResponseWriter, request *http.Request) {
		defer rest.save()
		id := request.URL.Query().Get("game_id")
		gameId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		game, ok := rest.games.Load(gameId)
		if !ok {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		c := request.URL.Query().Get("col")
		r := request.URL.Query().Get("row")
		col, err := strconv.ParseInt(c, 10, 64)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		row, err := strconv.ParseInt(r, 10, 64)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		sign, err := game.(*service.TicTacToe).Move(int(row), int(col))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}
		if sign != service.None {
			writer.Write([]byte(fmt.Sprintf("%d won!", sign)))
			return
		}
		json.NewEncoder(writer).Encode(game.(*service.TicTacToe).GetField())
	})

	rest.server = &http.Server{
		Addr:    listen,
		Handler: router,
	}

	rest.restore()

	return rest, nil
}

func (r *Rest) restore() error {
	res := map[int64]struct {
		Field   [][]service.Sign
		Current service.Sign
	}{}

	f, err := os.Open("games.json")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&res); err != nil {
		return err
	}

	for k, v := range res {
		fields, _ := service.NewTicTacToeFromField(v.Field, v.Current)
		r.games.Store(k, fields)
	}
	return nil
}
func (r *Rest) save() error {
	res := map[int64]struct {
		Field   [][]service.Sign
		Current service.Sign
	}{}

	r.games.Range(func(key, value any) bool {
		res[key.(int64)] = struct {
			Field   [][]service.Sign
			Current service.Sign
		}{
			Field:   value.(*service.TicTacToe).GetField(),
			Current: value.(*service.TicTacToe).Current(),
		}
		return true
	})

	f, err := os.Create("games.json")
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(res)
}

func (r *Rest) Run() error {
	if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (r *Rest) Close() error {
	return r.server.Shutdown(context.Background())
}
