package usecase

import (
	"strings"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/domain"
	"github.com/betorvs/sensubot/utils"
	"google.golang.org/api/chat/v1"
)

// ParseGoogleChatCommand func
func ParseGoogleChatCommand(data *chat.DeprecatedEvent) (string, error) {
	textRequest := strings.Split(strings.ToLower(data.Message.Text), " ")
	if len(textRequest) > 2 {
		go sendResultGoogleChat(data)
	}
	return "Accepted", nil
}

func sendResultGoogleChat(data *chat.DeprecatedEvent) {
	logLocal := config.GetLogger()
	role := "user"
	logLocal.Debug(config.Values.GoogleChatAdminList)
	if len(config.Values.GoogleChatAdminList) != 0 && utils.StringInSlice(data.Message.Sender.Name, config.Values.GoogleChatAdminList) {
		logLocal.Debug(config.Values.GoogleChatAdminList)
		role = "admin"
	}

	message := requestToBot(data.Message.ArgumentText, role, data.Message.Sender.DisplayName)

	repo := domain.GetGoogleChatRepository()

	// var sensuURL string
	// if config.Values.SensuURL != "Absent" {
	// 	sensuURL = config.Values.SensuURL
	// }
	err := repo.SendMessageGoogleChat(data, utils.Trim(message, 4000))
	if err != nil {
		logLocal.Error(err)
	}
	logLocal.Info("message send to google chat")

}
