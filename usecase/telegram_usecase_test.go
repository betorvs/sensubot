package usecase

import (
	"strings"
	"sync"
	"testing"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/domain"
)

var (
	TelegramMessageCalls   int
	TelegramMessageCallsMu sync.Mutex
)

type TelegramRepositoryMock struct {
}

func (repo TelegramRepositoryMock) SendTelegramMessage(reqBody []byte) error {
	TelegramMessageCallsMu.Lock()
	defer TelegramMessageCallsMu.Unlock()
	TelegramMessageCalls++
	return nil
}

func TestParseTelegramCommand(t *testing.T) {
	repo := TelegramRepositoryMock{}
	appcontext.Current.Add(appcontext.TelegramRepository, repo)
	message := new(domain.WebhookBody)
	message.Message.Text = "/help"
	message.Message.Chat.ID = int64(12345678)
	response, _ := ParseTelegramCommand(message)
	if !strings.Contains(response, "OK") {
		t.Fatalf("Invalid 1.1 TestParseTelegramCommand %s", response)
	}
}

func TestSendResultTelegram(t *testing.T) {
	repo := TelegramRepositoryMock{}
	appcontext.Current.Add(appcontext.TelegramRepository, repo)
	sendResultTelegram("gel all namespaces", int64(123456))
	expected := 1
	if TelegramMessageCalls != expected {
		t.Fatalf("Invalid 2.1 TestSendResultTelegram %d", TelegramMessageCalls)
	}
}
