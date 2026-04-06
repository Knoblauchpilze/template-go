package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Knoblauchpilze/backend-toolkit/pkg/config"
	"github.com/Knoblauchpilze/backend-toolkit/pkg/db"
	"github.com/Knoblauchpilze/backend-toolkit/pkg/logger"
	"github.com/Knoblauchpilze/backend-toolkit/pkg/process"
	"github.com/Knoblauchpilze/backend-toolkit/pkg/server"
	"github.com/Knoblauchpilze/template-go/cmd/app/internal"
	"github.com/Knoblauchpilze/template-go/internal/controller"
)

func determineConfigName() string {
	if len(os.Args) < 2 {
		return "config-prod.yml"
	}

	return os.Args[1]
}

func main() {
	log := logger.New(os.Stdout)

	conf, err := config.Load(determineConfigName(), internal.DefaultConfig())
	if err != nil {
		log.Error("Failed to load configuration", slog.Any("error", err))
		os.Exit(1)
	}

	conn, err := db.New(context.Background(), conf.Database)
	if err != nil {
		log.Error("Failed to create db connection", slog.Any("error", err))
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	s := server.NewWithLogger(conf.Server, log)

	for _, route := range controller.HealthCheckEndpoints(conn) {
		if err := s.AddRoute(route); err != nil {
			log.Error("Failed to register route", slog.String("route", route.Path()), slog.Any("error", err))
			os.Exit(1)
		}
	}

	wait, err := process.StartWithSignalHandler(context.Background(), s)
	if err != nil {
		log.Error("Failed to start server", slog.Any("error", err))
		os.Exit(1)
	}

	err = wait()
	if err != nil {
		log.Error("Error while serving", slog.Any("error", err))
		os.Exit(1)
	}
}
