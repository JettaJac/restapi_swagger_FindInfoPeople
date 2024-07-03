package app

import (
	// "database/sql"
	// "encoding/json"
	"fmt"
	// "github.com/go-openapi/spec"
	"log/slog"
	// swapi "main/generated"

	"main/internal/app/api"
	"main/internal/config"
	"main/pkg/lib/logger"
	// "main/internal/server"
	"main/internal/service/people"
	"main/internal/storage/postgre"
	// "net/http"
	// "os"
)

type App struct {
	ApiSrv *apiapp.App
}

func NewApp(config *config.Config, log *slog.Logger) (*App, error) {
	// DB

	const op = "internal/app.New"

	storage, err := postgre.New(config.NameDataBase /*config.Postgres*/)
	if err != nil {
		log.Error("app.NewApp: Failed to init storage", op, sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Сonnected to the database", slog.String("env", config.Env))
	defer storage.CloseDB()

	// timeTrackerService := Tracker
	peopleService := people.New(log, storage /*, storage*/)
	log.Info("Starting app")

	return &App{
		ApiSrv: apiapp.New(log, config, *peopleService), // * не было
	}, nil
}

// type Storage struct {
// 	db *sql.DB
// }

// func New(storagePath string) (*Storage, error) {
// 	const op = "storage.posgre.New"

// 	db, err := sql.Open("postgres", storagePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %s", op, err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		return nil, fmt.Errorf("%s.Ping: %w", op, err)
// 	}

// 	// возможно здесь запустить миграции как в п.посгре

// 	return &Storage{db: db}, nil
// }

// var _ swapi.ServerInterface = (*Storage)(nil)

// // sendPetStoreError wraps sending of an error in the Error format, and
// // handling the failure to marshal that.
// func sendPetStoreError(w http.ResponseWriter, code int, message string) { // обработака ошибок
// 	petErr := swapi.Error{
// 		Code:    int32(code),
// 		Message: message,
// 	}
// 	w.WriteHeader(code)
// 	_ = json.NewEncoder(w).Encode(petErr)
// }

// // Хендлер

// func (p *Storage) GetInfo(w http.ResponseWriter, r *http.Request, params swapi.GetInfoParams) {
// 	result := []string{"Все ", "работает"}

// 	// if params.Limit != nil {
// 	// 		l := int(*params.Limit)
// 	// 		if len(result) >= l {
// 	// 			// We're at the limit
// 	// 			break
// 	// 		}
// 	// 	}

// 	w.WriteHeader(http.StatusOK)
// 	_ = json.NewEncoder(w).Encode(result)
// }
