package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Knoblauchpilze/backend-toolkit/pkg/db"
	"github.com/Knoblauchpilze/backend-toolkit/pkg/db/postgresql"
	"github.com/labstack/echo/v5"
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

func generateTestEchoContextFromRequest(req *http.Request) (*echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	rw := httptest.NewRecorder()

	ctx := e.NewContext(req, rw)
	return ctx, rw
}
