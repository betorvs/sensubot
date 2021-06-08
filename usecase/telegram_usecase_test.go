package usecase

import (
	"testing"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/domain"
	localtest "github.com/betorvs/sensubot/test"
	"github.com/stretchr/testify/assert"
)

func TestSendResultTelegram(t *testing.T) {
	appcontext.Current.Add(appcontext.TelegramRepository, localtest.InitTelegramMock)
	appcontext.Current.Add(appcontext.Logger, localtest.InitMockLogger)
	sendResultTelegram("@sensubot gel namespaces", "Test User", int64(123123123), int64(123456))
	expected := 1
	assert.Equal(t, expected, localtest.TelegramCalls)
}

func TestParseTelegramCommand(t *testing.T) {
	appcontext.Current.Add(appcontext.TelegramRepository, localtest.InitTelegramMock)
	message := new(domain.WebhookBody)
	message.Message.Text = "/help"
	message.Message.Chat.ID = 12345678
	response, _ := ParseTelegramCommand(message)
	assert.Equal(t, "Accepted", response)
}
