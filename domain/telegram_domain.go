package domain

import "github.com/betorvs/sensubot/appcontext"

// WebhookBody struct
// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type WebhookBody struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int64  `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID                          int64  `json:"id"`
			Title                       string `json:"title"`
			Type                        string `json:"type"`
			AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
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
