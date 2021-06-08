package domain

import (
	"github.com/betorvs/sensubot/appcontext"
	"google.golang.org/api/chat/v1"
)

// GoogleChatRepository interface
type GoogleChatRepository interface {
	appcontext.Component
	// SendMessageGoogleChat requires *chat.DeprecatedEvent plus string and return error
	SendMessageGoogleChat(event *chat.DeprecatedEvent, content string) error
}

// GetGoogleChatRepository func return GoogleChatRepository interface
func GetGoogleChatRepository() GoogleChatRepository {
	return appcontext.Current.Get(appcontext.GoogleChatRepository).(GoogleChatRepository)
}
