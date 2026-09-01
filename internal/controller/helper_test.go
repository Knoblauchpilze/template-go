package controller

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Knoblauchpilze/backend-toolkit/pkg/db"
	"github.com/Knoblauchpilze/backend-toolkit/pkg/db/postgresql"
	"github.com/Knoblauchpilze/backend-toolkit/pkg/rest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var dbTestConfig = postgresql.NewConfigForLocalhost("db_template_service", "template_service_manager", "manager_password")

func newTestConnection(t *testing.T) db.Connection {
	t.Helper()

	conn, err := db.New(t.Context(), dbTestConfig)
	require.NoError(t, err, "Actual err: %v", err)

	t.Cleanup(func() {
		conn.Close(t.Context())
	})

	return conn
}

func generateTestRequest(
	t *testing.T,
	method string,
	modifiers ...func(*testing.T, *http.Request),
) *http.Request {
	t.Helper()

	ctx := rest.WithContextLogger(t.Context(), slog.Default())
	req := httptest.NewRequestWithContext(ctx, method, "/", nil)

	for _, modifier := range modifiers {
		modifier(t, req)
	}

	return req
}

func createTestGinRouter(
	t *testing.T,
	method string,
	path string,
	handler gin.HandlerFunc,
	middlewares ...gin.HandlerFunc,
) *gin.Engine {
	t.Helper()

	r := gin.New()

	for _, middleware := range middlewares {
		r.Use(middleware)
	}

	r.Handle(method, path, handler)

	return r
}

func decodeResponseBody[T any](t *testing.T, w *httptest.ResponseRecorder) T {
	t.Helper()

	var responseBody T

	rawBody, err := io.ReadAll(w.Result().Body)
	require.NoError(t, err, "Actual err: %v", err)

	err = json.Unmarshal(rawBody, &responseBody)
	require.NoError(t, err, "Actual err: %v", err)

	return responseBody
}
