package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostGChatEvents(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sensubot/v1/gchat")

	// Assertions
	if assert.NoError(t, GChatEvents(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}
