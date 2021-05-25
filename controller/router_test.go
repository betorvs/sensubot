package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/sensubot/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMapRoutesDefault(t *testing.T) {
	// Setup
	e := echo.New()
	MapRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sensubot/v1")
	f := func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, nil)
	}

	if assert.NoError(t, f(c)) {
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
}

func TestMapRoutesMetrics(t *testing.T) {
	// Setup
	config.Values.UsePrometheus = true
	e := echo.New()
	MapRoutes(e)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/metrics")

	assert.Equal(t, http.StatusOK, rec.Code)
}
