package apiapp

import (
	"context"
	"fmt"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"log/slog"
	swapi "main/generated"
	"main/internal/config"
	"main/internal/server"
	"main/internal/service/people"
	"net/http"
)

type App struct {
	log    *slog.Logger
	server *server.Server
	info   people.Info
}

func New(log *slog.Logger, config *config.Config, PeopleProvider people.Info) *App {
	fmt.Println("ttt")
	apiServer := server.NewServer(config, log, PeopleProvider)

	log.Info("Starting server", slog.String("address", config.HTTPServer.Address))

	return &App{
		log:    log,
		server: apiServer,
		info:   PeopleProvider,
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

	// a.server.ConfigureRouter(router)

	// Регистрация обработчиков Swagger
	z := swapi.HandlerFromMux(a.server, router)
	_ = z

	// Применение валидатора запросов
	h := middleware.OapiRequestValidator(swagger)(router)
	_ = h
	a.server.Server.Handler = z

	// router.Handle("/swagger", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	m := []byte("swagger.MarshalJSON()")
	// 	w.Write(m)
	// }))

	// s := &http.Server{
	// 	Handler: z,
	// 	Addr:    "0.0.0.0:8080",
	// }

	a.server.ConfigureRouter(router)

	// _ = s
	// Запуск сервера|
	// !!! Проверить через какой роутер запускаеться
	a.log.Info("Server is starting++++", slog.String("address", "a.config.HTTPServer.Address) --- подумать надо ли"))
	// return http.ListenAndServe(":8080", router) // !!! изменить на корректный адрес
	// return http.ListenAndServe(":8080", a.server) // !!! изменить на корректный адрес
	// return a.server.Server.ListenAndServe()
	// return s.ListenAndServe()
	return http.ListenAndServe(":8080", h)
	// return http.ListenAndServe(":8080", a.server.Router)
	// return http.ListenAndServe(":8080", a.server.Server.Handler)
	// return nil
}

// Stop stops the gRPC server.
func (a *App) Stop(ctx context.Context) {
	const op = "apiapp.Stop"
	log := a.log.With(slog.String("op", op))

	log.Info("stopping API server", slog.String("API address", a.server.Server.Addr))
	// a.info.
	a.server.Server.Shutdown(ctx)
}
