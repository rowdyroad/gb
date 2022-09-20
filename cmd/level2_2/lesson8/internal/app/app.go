package app

import "gb/cmd/level2_2/lesson8/internal/delivery/rest"

type Config struct {
	Size        int
	Player1name string
	Player2name string
	Listen      string // 0.0.0.0:80
}

type App struct {
	config Config
	rest   *rest.Rest
}

func NewApp(config Config) (*App, error) {
	a := &App{
		config: config,
	}
	var err error
	a.rest, err = rest.NewRest(config.Listen, config.Size)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	return a.rest.Run()
}

func (a *App) Close() error {
	return a.rest.Close()
}
