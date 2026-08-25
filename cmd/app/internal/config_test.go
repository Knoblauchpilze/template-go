package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_DefaultConfig(t *testing.T) {
	t.Run("provides default REST configuration", func(t *testing.T) {
		config := DefaultConfig()

		assert.Equal(t, "/v1", config.Server.BasePath)
		assert.Equal(t, uint16(80), config.Server.Port)
	})

	t.Run("provides default database connection", func(t *testing.T) {
		config := DefaultConfig()

		assert.Equal(t, "172.17.0.1", config.Database.Host)
		assert.Equal(t, "db_template_service", config.Database.Database)
		assert.Equal(t, "template_service_manager", config.Database.User)
	})

	t.Run("does not set database password", func(t *testing.T) {
		config := DefaultConfig()

		assert.Equal(t, "comes-from-the-environment", config.Database.Password)
	})
}
