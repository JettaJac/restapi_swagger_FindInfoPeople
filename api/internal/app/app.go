package app

import (
	// "database/sql"
	// "encoding/json"
	"fmt"
	"log/slog"
	"main/internal/app/api"
	"main/internal/config"
	"main/internal/service/people"
	"main/internal/storage/postgre"
	"main/pkg/lib/logger"
)

type App struct {
	ApiSrv *apiapp.App
}

func NewApp(config *config.Config, log *slog.Logger) (*App, error) {

	const op = "internal/app.New"

	storage, err := postgre.New(config.Postgres)
	if err != nil {
		log.Error("app.NewApp: Failed to init storage", op, sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Сonnected to the database", slog.String("env", config.Env))
	// defer storage.CloseDB()

	// // timeTrackerService := Tracker
	// peopleService := people.New(log, storage /*, storage*/)
	// log.Info("Starting app")
	// // app := server.NewServer(cfg)

	authService := people.New(log, storage) //переделать на 1 стораж
	log.Info("Starting app")
	_ = authService

	apiApp := apiapp.New(log, config, *authService) // * не было
	log.Info("Starting APIserver")                  //  не уверена что правильно

	return &App{
		ApiSrv: apiApp,
	}, nil
}
