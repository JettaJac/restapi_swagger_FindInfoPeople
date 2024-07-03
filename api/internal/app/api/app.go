package apiapp

import (
	"fmt"
	// middleware "github.com/oapi-codegen/nethttp-middleware"
	"context"
	"log/slog"
	swapi "main/generated"
	"main/internal/config"
	"main/internal/server"
	"net/http"
)

type App struct {
	// config *config.Config // !!! А нужен ли здесь конфиг
	log    *slog.Logger
	server *server.Server
}

// type People interface {
// 	GetInfo(http.ResponseWriter, *http.Request, swapi.GetInfoParams)
// }

func New(log *slog.Logger, config *config.Config /*, PeopleProvider people.People*/) *App {
	fmt.Println("ttt")
	apiServer := server.NewServer(config, log /*, swagger*/)

	log.Info("Starting server", slog.String("address", config.HTTPServer.Address))

	return &App{
		// config: config,
		log:    log,
		server: apiServer,
	}

}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	// !!!Логгер уже инициализирован ранее

	// Загрузка спецификации Swagger
	swagger, err := swapi.GetSwagger()
	if err != nil {
		a.log.Error("Error loading swagger spec" /*, slog.Error(err)*/)
		return err
	}
	swagger.Servers = nil // Удаление серверов из спецификации, если они там есть

	// Создание маршрутизатора
	router := http.NewServeMux()

	// Регистрация пользовательского обработчика GetInfo //!!! Вынест в отдельную функцию
	// router.Handle("/info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// Создание и заполнение структуры параметров
	// 	params := swapi.GetInfoParams{
	// 		PassportSerie:  7777777,
	// 		PassportNumber: 3333333,
	// 	}
	// 	a.server.GetInfo(w, r, params)
	// }))

	a.server.ConfigureRouter(router)

	// Регистрация обработчиков Swagger
	swapi.HandlerFromMux(a.server, router)

	// Применение валидатора запросов
	// h := middleware.OapiRequestValidator(swagger)(router)

	// Запуск сервера
	a.log.Info("Server is starting", slog.String("address", "a.config.HTTPServer.Address) --- подумать надо ли"))
	return http.ListenAndServe(":8080", router) // !!! изменить на корректный адрес
}

// Stop stops the gRPC server.
func (a *App) Stop(ctx context.Context) {
	const op = "apiapp.Stop"
	log := a.log.With(slog.String("op", op))
	log.Info("stopping API server", slog.String("API address", a.server.Server.Addr))

	a.server.Server.Shutdown(ctx)
}
