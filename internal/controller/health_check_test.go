package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Knoblauchpilze/backend-toolkit/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIT_HealthcheckController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("returns ok when connection is healthy", func(t *testing.T) {
		conn := newTestConnection(t)
		handler := createServiceAwareHttpHandler(healthcheck, conn)

		r := createTestGinRouter(t, http.MethodGet, "/", handler)

		req := generateTestRequest(t, http.MethodGet)
		rw := httptest.NewRecorder()
		r.ServeHTTP(rw, req)

		assert.Equal(t, http.StatusOK, rw.Code)
		actual := decodeResponseBody[string](t, rw)
		assert.Equal(t, "OK", actual)
	})

	t.Run("returns error when connection is not healthy", func(t *testing.T) {
		conn := newTestConnection(t)
		conn.Close(t.Context())
		handler := createServiceAwareHttpHandler(healthcheck, conn)

		r := createTestGinRouter(t, http.MethodGet, "/", handler)

		req := generateTestRequest(t, http.MethodGet)
		rw := httptest.NewRecorder()
		r.ServeHTTP(rw, req)

		assert.Equal(t, http.StatusServiceUnavailable, rw.Code)
		actual := decodeResponseBody[errors.ErrorWithCode](t, rw)
		assert.Equal(t, errors.ErrorCode(100), actual.Code)
		assert.Equal(t, "an unexpected error occurred", actual.Message)
	})
}
