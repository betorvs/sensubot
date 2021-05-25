package usecase

import (
	"testing"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/domain"
	localtest "github.com/betorvs/sensubot/test"
	"github.com/stretchr/testify/assert"
)

func TestParseTelegramCommand(t *testing.T) {
	appcontext.Current.Add(appcontext.TelegramRepository, localtest.InitTelegramMock)
	message := new(domain.WebhookBody)
	message.Message.Text = "/help"
	message.Message.Chat.ID = int64(12345678)
	response, _ := ParseTelegramCommand(message)
	assert.Equal(t, "OK", response)
}

func TestSendResultTelegram(t *testing.T) {
	appcontext.Current.Add(appcontext.TelegramRepository, localtest.InitTelegramMock)
	sendResultTelegram("gel all namespaces", int64(123456))
	expected := 1
	assert.Equal(t, expected, localtest.TelegramCalls)
}
