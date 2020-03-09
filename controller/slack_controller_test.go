package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostSlackEvents(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sensubot/v1/slack")

	// Assertions
	if assert.NoError(t, SlackEvents(c)) {
		assert.Equal(t, http.StatusNotImplemented, rec.Code)
	}
}
