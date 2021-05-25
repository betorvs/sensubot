package usecase

import (
	"fmt"
	"testing"

	"github.com/betorvs/sensubot/appcontext"
	localtest "github.com/betorvs/sensubot/test"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestValidateBot(t *testing.T) {
	longString := string("2aeccc9c03b36fea59ebec69")
	wrongLongString := string("1656ca38de0749bb8620814f")
	bodyString := string("body")
	timestamp := "1580475458"
	baseString := fmt.Sprintf("v0:%s:%s", timestamp, bodyString)
	signature := "v0=38918cc13bf5e8b9ad7d2b7d85595d5d19af21783b86dbcfe4716422277ae404"
	assert.True(t, ValidateBot(signature, baseString, longString))
	assert.False(t, ValidateBot(signature, baseString, wrongLongString))
}

func TestParseSlashCommand(t *testing.T) {
	test1 := slack.SlashCommand{
		Text: "help",
	}
	response, _ := ParseSlashCommand(&test1)
	assert.Contains(t, response.Text, "Please use")
}

func TestSendResultSlack(t *testing.T) {
	appcontext.Current.Add(appcontext.SlackRepository, localtest.InitSlackMock)
	sendResultSlack("gel all namespaces", "USER123", "CHANNEL123")
	expected := 1
	assert.Equal(t, expected, localtest.SlackCalls)

}
