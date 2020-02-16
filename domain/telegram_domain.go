package domain

import "github.com/betorvs/sensubot/appcontext"

// WebhookBody struct
// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type WebhookBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// SendMessageReqBody struct
// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage
type SendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// TelegramRepository interface
type TelegramRepository interface {
	appcontext.Component
	// SendTelegramMessage requires chatID int64, reqBody []byte and return error
	SendTelegramMessage(reqBody []byte) error
}

// GetTelegramRepository func return TelegramRepository interface
func GetTelegramRepository() TelegramRepository {
	return appcontext.Current.Get(appcontext.TelegramRepository).(TelegramRepository)
}
