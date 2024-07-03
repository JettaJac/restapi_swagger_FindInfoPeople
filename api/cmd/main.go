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

// curl http://localhost:8080/info\?passportSerie=1234\&passportNumber=4663423

import (
	// middleware "github.com/oapi-codegen/nethttp-middleware"

	"context"
	"fmt"
	"log/slog"
	"main/internal/app"
	"main/internal/config"
	"main/pkg/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Start") //!!! Delete

	// initializing the configuration
	cfg := config.NewConfig()

	// logger initialization
	log := sl.SetupLogger(cfg.Env)
	log.Info("starting application", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// run application
	app, err := app.NewApp(cfg, log)
	if err != nil {
		log.Error("cannot create server", sl.Err(err))
	}
	// run server
	go app.ApiSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop
	log.Info("stopping signal", slog.String("signal", sign.String()))
	app.ApiSrv.Stop(context.Background())
	log.Info("application stopped")

}
