package main

import (
	"flag"

	"gb/cmd/level2_2/lesson8/internal/app"
)

func main() {
	var cfg app.Config
	flag.IntVar(&cfg.Size, "size", 3, "size of board")
	flag.StringVar(&cfg.Listen, "listen", "0.0.0.0:80", "http listen endpoint")
	flag.StringVar(&cfg.Player1name, "player1", "Player1", "")
	flag.StringVar(&cfg.Player2name, "player2", "Player2", "")
	flag.Parse()

	application, err := app.NewApp(cfg)
	if err != nil {
		panic(err)
	}
	defer application.Close()

	if err = application.Run(); err != nil {
		panic(err)
	}
}
