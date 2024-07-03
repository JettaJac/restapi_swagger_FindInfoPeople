package app

import (
	// "database/sql"
	// "encoding/json"
	"fmt"
	"log/slog"
	"main/internal/app/api"
	"main/internal/config"
	"main/internal/storage/postgre"
	"main/pkg/lib/logger"
)

type App struct {
	ApiSrv *apiapp.App
}

func NewApp(config *config.Config, log *slog.Logger) (*App, error) {

	const op = "internal/app.New"

	storage, err := postgre.New(config.Postgres /*config.Postgres*/)
	if err != nil {
		log.Error("app.NewApp: Failed to init storage", op, sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Сonnected to the database", slog.String("env", config.Env))
	defer storage.CloseDB()

	// // timeTrackerService := Tracker
	// peopleService := people.New(log, storage /*, storage*/)
	// log.Info("Starting app")
	// // app := server.NewServer(cfg)

	return &App{
		ApiSrv: apiapp.New(log, config /*, *peopleService*/), // * не было
	}, nil
}
