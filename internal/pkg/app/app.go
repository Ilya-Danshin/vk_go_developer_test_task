package app

import (
	"vk_go_develop_test_task/internal/app/bot"
	"vk_go_develop_test_task/internal/app/config"
)

type App struct {
	bot *bot.Bot
}

func New() (*App, error) {
	a := &App{}
	var err error

	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	a.bot, err = bot.New(cfg)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	err := a.bot.Run()
	if err != nil {
		return err
	}

	return nil
}
