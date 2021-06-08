package usecase

import (
	"testing"

	"github.com/betorvs/sensubot/appcontext"
	localtest "github.com/betorvs/sensubot/test"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/chat/v1"
	"google.golang.org/api/googleapi"
)

func TestParseGoogleChatCommand(t *testing.T) {
	appcontext.Current.Add(appcontext.GoogleChatRepository, localtest.InitGoogleChatMock)
	appcontext.Current.Add(appcontext.Logger, localtest.InitMockLogger)
	data := chat.DeprecatedEvent{
		Action:                    &chat.FormAction{},
		ConfigCompleteRedirectUrl: "",
		EventTime:                 "",
		Message: &chat.Message{
			ActionResponse:  &chat.ActionResponse{},
			Annotations:     []*chat.Annotation{},
			ArgumentText:    "",
			Attachment:      []*chat.Attachment{},
			Cards:           []*chat.Card{},
			CreateTime:      "",
			FallbackText:    "",
			Name:            "",
			PreviewText:     "",
			Sender:          &chat.User{},
			SlashCommand:    &chat.SlashCommand{},
			Space:           &chat.Space{},
			Text:            "get",
			Thread:          &chat.Thread{},
			ServerResponse:  googleapi.ServerResponse{},
			ForceSendFields: []string{},
			NullFields:      []string{},
		},
		Space:           &chat.Space{},
		ThreadKey:       "",
		Token:           "",
		Type:            "",
		User:            &chat.User{},
		ForceSendFields: []string{},
		NullFields:      []string{},
	}
	res1, err1 := ParseGoogleChatCommand(&data)
	assert.NoError(t, err1)
	assert.Equal(t, "Accepted", res1)

}
