package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/domain"
	"github.com/nlopes/slack"
)

// ValidateBot func to validate auth from Slack Bot
func ValidateBot(signing string, message string, mysigning string) bool {
	mac := hmac.New(sha256.New, []byte(mysigning))
	if _, err := mac.Write([]byte(message)); err != nil {
		log.Printf("[ERROR] ValidateBot mac.Write(%v) failed\n", message)
		return false
	}
	calculatedMAC := "v0=" + hex.EncodeToString(mac.Sum(nil))
	// log.Printf("[INFO] ValidateBot %s : %s ", signing, calculatedMAC)
	return hmac.Equal([]byte(calculatedMAC), []byte(signing))
}

// ParseSlashCommand func to parse Slack Text field and try to answer something useful to user
func ParseSlashCommand(data *slack.SlashCommand) (slack.Msg, error) {
	var res slack.Msg
	testRequest := strings.Split(strings.ToLower(data.Text), " ")
	if len(testRequest) == 1 {
		// verb specialWord|resource resouce|name specialWord namespace_name specialWord entity_name
		text := fmt.Sprintf("Invalid request received: %s \nPlease use: %s VERB RESOURCE NAME NAMESPACE", data.Text, config.SlackSlashCommand)
		message := slack.Msg{
			Text: text}
		res = message
	} else {
		go sendResultSlack(data.Text, data.UserID, data.ChannelID)
		message := slack.Msg{
			ResponseType: "in_channel",
			Text:         "Request Accepted"}
		res = message

	}
	return res, nil
}

func sendResultSlack(text string, userid string, channel string) {
	message := requestSensu(text)

	// Send a message to user in slack
	slack := domain.GetSlackRepository()
	var errSlack error
	if len([]rune(message)) > 3000 {
		errSlack = slack.EphemeralFileMessage(channel, userid, message, "attach")
	} else {
		errSlack = slack.EphemeralMessage(channel, userid, message)
	}
	if errSlack != nil {
		log.Printf("[ERROR] sendResultSlack %s", errSlack)
	}
}
