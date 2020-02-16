package usecase

import (
	"fmt"
	"strings"
	"testing"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/nlopes/slack"
)

var (
	slackEphemeralMessageCalls     int
	slackEphemeralFileMessageCalls int
)

func TestValidateBot(t *testing.T) {
	longString := string("2aeccc9c03b36fea59ebec69")
	wrongLongString := string("1656ca38de0749bb8620814f")
	bodyString := string("body")
	timestamp := "1580475458"
	baseString := fmt.Sprintf("v0:%s:%s", timestamp, bodyString)
	signature := "v0=38918cc13bf5e8b9ad7d2b7d85595d5d19af21783b86dbcfe4716422277ae404"
	if ValidateBot(signature, baseString, longString) != true {
		t.Fatalf("Invalid 1.1 testValidateBot %s, %s", signature, longString)
	}
	if ValidateBot(signature, baseString, wrongLongString) == true {
		t.Fatalf("Invalid 1.2 testValidateBot %s, %s", signature, wrongLongString)
	}
}

func TestParseSlashCommand(t *testing.T) {
	test1 := slack.SlashCommand{
		Text: "help",
	}
	response, _ := ParseSlashCommand(&test1)
	if !strings.Contains(response.Text, "help") {
		t.Fatalf("Invalid 2.1 TestParseSlashCommand %s", response.Text)
	}
}

type SlackRepositoryMock struct {
}

func (repo SlackRepositoryMock) EphemeralMessage(channel string, user string, message string) error {
	slackEphemeralMessageCalls++
	return nil
}

func (repo SlackRepositoryMock) EphemeralFileMessage(channel string, user string, message string, title string) error {
	slackEphemeralFileMessageCalls++
	return nil
}

func TestSendResultSlack(t *testing.T) {
	repo := SlackRepositoryMock{}
	appcontext.Current.Add(appcontext.SlackRepository, repo)
	sendResultSlack("gel all namespaces", "USER123", "CHANNEL123")
	expected := 1
	if slackEphemeralMessageCalls != expected {
		t.Fatalf("Invalid 3.1 TestSendResultSlack %d", expected)
	}

}
