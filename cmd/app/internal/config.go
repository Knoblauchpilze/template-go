package internal

import (
	"time"

	"github.com/Knoblauchpilze/backend-toolkit/pkg/db/postgresql"
	"github.com/Knoblauchpilze/backend-toolkit/pkg/server"
)

type Configuration struct {
	Server   server.Config
	Database postgresql.Config
}

func DefaultConfig() Configuration {
	const defaultDatabaseName = "db_template_service"
	const defaultDatabaseUser = "template_service_manager"

	return Configuration{
		Server: server.Config{
			BasePath:        "/v1",
			Port:            uint16(80),
			ShutdownTimeout: 5 * time.Second,
		},
		Database: postgresql.NewConfigForDockerContainer(
			defaultDatabaseName,
			defaultDatabaseUser,
			"comes-from-the-environment",
		),
	}
}
