package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIT_HealthcheckController(t *testing.T) {
	conn := newTestConnection(t)

	req := httptest.NewRequest(http.MethodGet, "/healtcheck", nil)
	ctx, rw := generateTestEchoContextFromRequest(req)

	err := healthcheck(ctx, conn)
	require.NoError(t, err, "Actual err: %v", err)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Equal(t, "\"OK\"\n", rw.Body.String())
}

func TestIT_HealthcheckController_WhenConnectionClosed_ExpectServiceUnavailable(t *testing.T) {
	conn := newTestConnection(t)
	conn.Close(t.Context())

	req := httptest.NewRequest(http.MethodGet, "/healtcheck", nil)
	ctx, rw := generateTestEchoContextFromRequest(req)

	err := healthcheck(ctx, conn)
	require.NoError(t, err, "Actual err: %v", err)

	assert.Equal(t, http.StatusServiceUnavailable, rw.Code)
	expectedResponse := `
	{
		"code": 100,
		"message": "an unexpected error occurred"
	}`
	assert.JSONEq(t, expectedResponse, rw.Body.String())
}
