package main

import (
	"context"
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
	log := logger.New(logger.NewPrettyWriter(os.Stdout))

	conf, err := config.Load(determineConfigName(), internal.DefaultConfig())
	if err != nil {
		log.Errorf("Failed to create db connection: %v", err)
		os.Exit(1)
	}

	conn, err := db.New(context.Background(), conf.Database)
	if err != nil {
		log.Errorf("Failed to create db connection: %v", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	s := server.NewWithLogger(conf.Server, log)

	for _, route := range controller.HealthCheckEndpoints(conn) {
		if err := s.AddRoute(route); err != nil {
			log.Errorf("Failed to register route %v: %v", route.Path(), err)
			os.Exit(1)
		}
	}

	wait, err := process.StartWithSignalHandler(context.Background(), s)
	if err != nil {
		log.Errorf("Failed to start server: %v", err)
		os.Exit(1)
	}

	err = wait()
	if err != nil {
		log.Errorf("Error while serving: %v", err)
		os.Exit(1)
	}
}
