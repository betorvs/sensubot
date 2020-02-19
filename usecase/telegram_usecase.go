package usecase

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/betorvs/sensubot/domain"
)

// ParseTelegramCommand func
func ParseTelegramCommand(data *domain.WebhookBody) (string, error) {
	textRequest := strings.Split(strings.ToLower(data.Message.Text), " ")
	if len(textRequest) > 2 {
		go sendResultTelegram(data.Message.Text, data.Message.Chat.ID)
	}
	return "OK", nil
}

func sendResultTelegram(text string, chatID int64) {
	message := requestSensu(text)

	reqBody := &domain.SendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(&reqBody)
	if err != nil {
		log.Printf("[ERROR] sendResultTelegram Marshal %s", err)
	}
	telegram := domain.GetTelegramRepository()
	errTelegram := telegram.SendTelegramMessage(reqBytes)
	if errTelegram != nil {
		log.Printf("[ERROR] sendResultTelegram %s", errTelegram)
	}

}

// SendHelpMessage func is used to send a help message to user
func SendHelpMessage(chatID int64) error {
	message := "Use: get|execute|silence server check\n Like: get all checks\n Or: silence check http \n Or: get health\n More info: github.com/betorvs/sensubot"
	reqBody := &domain.SendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(&reqBody)
	if err != nil {
		log.Printf("[ERROR] sendHelpMessage Marshal %s", err)
		return err
	}
	telegram := domain.GetTelegramRepository()
	errTelegram := telegram.SendTelegramMessage(reqBytes)
	if errTelegram != nil {
		log.Printf("[ERROR] sendHelpMessage %s", errTelegram)
		return errTelegram
	}
	return nil
}
