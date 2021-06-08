package test

import (
	"github.com/betorvs/sensubot/appcontext"
	"google.golang.org/api/chat/v1"
)

var (
	// GoogleChatCalls int
	GoogleChatCalls int
)

// MockGoogleChatRepository struct
type MockGoogleChatRepository struct {
}

func (repo MockGoogleChatRepository) SendMessageGoogleChat(event *chat.DeprecatedEvent, content string) error {
	GoogleChatCalls++
	return nil
}

// InitGoogleChatMock returns a MockGoogleChatRepository interface
func InitGoogleChatMock() appcontext.Component {

	return MockGoogleChatRepository{}
}
