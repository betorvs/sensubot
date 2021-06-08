package usecase

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/domain"
	"github.com/betorvs/sensubot/utils"
)

// ParseTelegramCommand func
func ParseTelegramCommand(data *domain.WebhookBody) (string, error) {
	textRequest := strings.Split(strings.ToLower(data.Message.Text), " ")
	if len(textRequest) >= 2 {
		displayName := fmt.Sprintf("%s %s", data.Message.From.FirstName, data.Message.From.LastName)
		go sendResultTelegram(data.Message.Text, displayName, data.Message.From.ID, data.Message.Chat.ID)
	}
	return "Accepted", nil
}

func sendResultTelegram(text, displayName string, username, chatID int64) {
	logLocal := config.GetLogger()
	logLocal.Debug("starting sendResultTelegram")
	parsedText := text
	if strings.HasPrefix(parsedText, config.Values.TelegramBotName) {
		parsedText = strings.TrimLeft(parsedText, config.Values.TelegramBotName)
	}
	role := "user"
	if len(config.Values.TelegramAdminIDList) != 0 && utils.StringInSlice(fmt.Sprintf("%v", username), config.Values.TelegramAdminIDList) {
		role = "admin"
	}
	message := requestToBot(parsedText, role, displayName)
	if message == "" {
		message = "Empty response from Integration"
	}
	reqBody := &domain.SendMessageReqBody{
		ChatID: chatID,
		Text:   utils.Trim(message, 4000),
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(&reqBody)
	if err != nil {
		logLocal.Error("sendResultTelegram Marshal %s", err)
	}
	telegram := domain.GetTelegramRepository()
	errTelegram := telegram.SendTelegramMessage(reqBytes)
	if errTelegram != nil {
		logLocal.Errorf("sendResultTelegram %s", errTelegram)
	}

}

// SendHelpMessage func is used to send a help message to user
func SendHelpMessage(chatID int64) error {
	message := "Use: get|execute|silence server check\n Like: get all checks\n Or: silence check http \n Or: get health\n More info: github.com/betorvs/sensubot"
	reqBody := &domain.SendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}
	logLocal := config.GetLogger()
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(&reqBody)
	if err != nil {
		logLocal.Errorf("sendHelpMessage Marshal %s", err)
		return err
	}
	telegram := domain.GetTelegramRepository()
	errTelegram := telegram.SendTelegramMessage(reqBytes)
	if errTelegram != nil {
		logLocal.Errorf("sendHelpMessage %s", errTelegram)
		return errTelegram
	}
	return nil
}
