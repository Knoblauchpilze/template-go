package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_DefaultConfig_DefinesCorrectRestConfiguration(t *testing.T) {
	config := DefaultConfig()

	assert.Equal(t, "/v1", config.Server.BasePath)
	assert.Equal(t, uint16(80), config.Server.Port)
}

func TestUnit_DefaultConfig_SetsExpectedDbConnection(t *testing.T) {
	config := DefaultConfig()

	assert.Equal(t, "172.17.0.1", config.Database.Host)
	assert.Equal(t, "db_products", config.Database.Database)
	assert.Equal(t, "products_manager", config.Database.User)
}

func TestUnit_DefaultConfig_DoesNotSetDbPassword(t *testing.T) {
	config := DefaultConfig()

	assert.Equal(t, "comes-from-the-environment", config.Database.Password)
}
