package main

// go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=swagger/swagger.yaml ../../api.yaml
// go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=swagger/swagger.yaml api/apigenerate/cfgswagger2.yaml
// oapi-codegen generate server,types -g api/apigenerate
// oapi-codegen generate server,types -g apigenerate ../swagger/swagger.yaml
// oapi-codegen --config api/apigenerate/cfgswagger2.yaml swagger/swagger.yaml
// из папки api oapi-codegen --config config/cfgswagger.yaml ../swagger/swagger.yaml
// oapi-codegen --config config/cfgswagger.yaml --output ./apigenerate/ ../swagger/swagger.yaml
// oapi-codegen generate server,types --output-dir apigenerate --output-file api.gen.go ../swagger/swagger.yaml
// oapi-codegen --config config/cfgswagger.yaml ../swagger/swagger.yaml  - done

import (
	"fmt"
	// _ "github.com/oapi-codegen/oapi-codegen"
	// _ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	// _ "github.com/oapi-codegen/oapi-codegen/v2"

	middleware "github.com/oapi-codegen/nethttp-middleware"

	"log/slog"

	swapi "main/generated"
	// "main/internal/app"
	"main/internal/config"
	"main/internal/server"
	"main/pkg/lib/logger"

	"net"
	"net/http"
	"os"
	// "os/signal"
	// "syscall"
)

// !!! Transfer lib from from internal to pkg

func main() {
	fmt.Println("Start") //!!! Delete
	// // create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	// server := api.NewServer()

	// r := http.NewServeMux()

	// // get an `http.Handler` that we can use
	// h := api.HandlerFromMux(server, r)

	// s := &http.Server{
	// 	Handler: h,
	// 	Addr:    "0.0.0.0:8080",
	// }

	// // And we serve HTTP until the world ends.
	// log.Fatal(s.ListenAndServe())

	// ------------------------------------------------------------------
	// !!!

	// Initializing the configuration
	cfg := config.NewConfig()

	// Logger initialization
	log := sl.SetupLogger(cfg.Env)
	log.Info("starting application", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// База данных
	/*
		// инициализация приложения
		application, err := app.NewApp(cfg, log)
		if err != nil {
			log.Error("The application is not initialized: ", sl.Err(err))
			return
		}

		// запуск самого сервера
		go application.ApiSrv.MustRun()

		// остановка сервера по сигналам
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
		sign := <-stop
		log.Info("stopping signal", slog.String("signal", sign.String()))
		application.ApiSrv.Stop()
		log.Info("application stopped")
	*/
	// ------------------------------------------------------------------

	// // port := flag.String("port", "8080", "Port for test HTTP server")
	// // flag.Parse()
	swagger, err := swapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// // Create an instance of our handler which satisfies the generated interface
	// petStore, _ := api.NewPetStore()
	// Storage, _ := api.New("fgchdj")
	// _ = Storage

	app := server.NewServer(cfg)

	r := http.NewServeMux()
	// We now register our petStore above as the handler for the interface
	swapi.HandlerFromMux(app, r)

	// // We now register our petStore above as the handler for the interface
	// swapi.HandlerFromMux(petStore, r)

	// // Use our validation middleware to check all requests against the
	// // OpenAPI schema.
	h := middleware.OapiRequestValidator(swagger)(r)

	s := &http.Server{
		Handler: h,
		// Handler: r,
		Addr: net.JoinHostPort("0.0.0.0", "8080"),
	}

	// // And we serve HTTP until the world ends.
	fmt.Println(s.ListenAndServe()) // тут д.б лог

}
