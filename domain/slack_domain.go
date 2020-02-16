package domain

import "github.com/betorvs/sensubot/appcontext"

// SlackRepository interface
type SlackRepository interface {
	appcontext.Component
	// EphemeralMessage func send a message using channelID, userID and textMessage and return an error
	EphemeralMessage(channel string, user string, message string) error
	// EphemeralMessage func send a message with attachment using channelID, userID and textMessage and return an error
	EphemeralFileMessage(channel string, user string, message string, title string) error
}

// GetSlackRepository func return SlackRepository interface
func GetSlackRepository() SlackRepository {
	return appcontext.Current.Get(appcontext.SlackRepository).(SlackRepository)
}
