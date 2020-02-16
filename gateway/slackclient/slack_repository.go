package slackclient

import (
	"fmt"
	"log"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
	"github.com/nlopes/slack"
)

// Slack struct for our slackbot
type Slack struct {
	Name  string
	Token string

	User   string
	UserID string

	Client *slack.Client
}

// EphemeralMessage func send a message using channelID, userID and textMessage and return an error
func (repo Slack) EphemeralMessage(channel string, user string, message string) error {
	attachment := slack.Attachment{
		Text: message,
	}
	if _, err := repo.Client.PostEphemeral(channel, user, slack.MsgOptionAttachments(attachment)); err != nil {
		return fmt.Errorf("[ERROR] EphemeralMessage failed to post message: %v", err)
	}
	return nil
}

// EphemeralFileMessage func send a message with attachment using channelID, userID and textMessage and return an error
func (repo Slack) EphemeralFileMessage(channel string, user string, message string, title string) error {
	if message == "" {
		message = "Empty"
	}
	params := slack.FileUploadParameters{
		Filename: "result.txt", Title: title, Content: message,
		Channels: []string{channel}}
	if _, err := repo.Client.UploadFile(params); err != nil {
		log.Printf("[ERROR] EphemeralFileMessage failed to post message with an file: %s", err)
	}
	return nil
}

// New func to initializate Slack Client
func New() (*Slack, error) {
	return &Slack{Client: slack.New(config.SlackToken), Token: config.SlackToken, Name: "Slack to Sensu Go"}, nil
}

func init() {
	if config.GetEnv("TESTRUN", "false") == "true" {
		return
	}
	SlackRepository, err := New()
	if err != nil {
		log.Println("[ERROR] Slack Repository not initiated")
	}
	appcontext.Current.Add(appcontext.SlackRepository, SlackRepository)
	if appcontext.Current.Count() != 0 {
		log.Println("[INFO] Slack Repository initiated")
	}

}
