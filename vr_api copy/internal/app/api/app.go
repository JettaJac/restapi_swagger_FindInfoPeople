package apiapp

import (
	// "database/sql"
	// "encoding/json"
	"fmt"
	// "github.com/go-openapi/spec"
	// api3 "github.com/oapi-codegen/oapi-codegen"
	"log/slog"
	swapi "main/generated"
	"main/internal/config"
	"main/internal/server"
	"main/internal/service/people"
	// "net/http"
	"os"
)

// // Генерация документации Swagger
// swaggerJSON, err := json.MarshalIndent(swagger, "", "  ")
//
//	if err != nil {
//	    // Обработка ошибки
//	}
//
// /
type App struct {
	config *config.Config // !!! А нужен ли здесь конфиг
	log    *slog.Logger
	server *server.Server
	// swagger *api3.Swagger
}

// type People interface {
// 	GetInfo(http.ResponseWriter, *http.Request, swapi.GetInfoParams)
// }

func New(log *slog.Logger, config *config.Config, PeopleProvider people.People) *App {
	swagger, err := swapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	server := server.NewServer(config /*, swagger*/)
	log.Info("Starting server", slog.String("address", config.Address))
	//  Для регистрации обработчиков API-маршрутов.
	// swapi.HandlerFromMux(PeopleProvider, server.Router) //!!! Возможно не здесь надо а в сервере
	return &App{
		config: config,
		log:    log,
		server: server,
		// swagger: swagger,
	}

}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "apiapp.Run"
	log := a.log.With(
		slog.String("op", op),
		// slog.Int("grpcPort", a.config.Address),подумать в конфигах какой формат инт или стринг
	)
	_ = log //!!!
	return nil
}

func (a *App) Stop() {
	// !!!
}
