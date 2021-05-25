package controller

import (
	"net/http"
	"strings"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/domain"
	"github.com/betorvs/sensubot/usecase"
	"github.com/labstack/echo/v4"
)

// TelegramEvents func receive an post from Telegram slash integration
func TelegramEvents(c echo.Context) error {
	if config.Values.TelegramToken == disabled {
		return c.JSON(http.StatusNotImplemented, nil)
	}
	event := new(domain.WebhookBody)
	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if strings.Contains(event.Message.Text, "/start") || strings.Contains(event.Message.Text, "/help") {
		err := usecase.SendHelpMessage(event.Message.Chat.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

	}
	res, err := usecase.ParseTelegramCommand(event)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusAccepted, res)
}
