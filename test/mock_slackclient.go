package test

import "github.com/betorvs/sensubot/appcontext"

var (
	// SlackCalls int
	SlackCalls int
)

// SlackMockRepository struct mock
type SlackMockRepository struct {
}

// EphemeralMessage func send a message using channelID, userID and textMessage and return an error
func (repo SlackMockRepository) EphemeralMessage(channel string, user string, message string) error {
	SlackCalls++
	return nil
}

// SendMessage func
func (repo SlackMockRepository) SendMessage(channel string, message, attach string) error {
	SlackCalls++
	return nil
}

// SendMessageWithImage func
func (repo SlackMockRepository) SendMessageWithImage(channel string, message, textImage, imageTitle, imageURL string) error {
	SlackCalls++
	return nil
}

// EphemeralFileMessage func send a message with attachment using channelID, userID and textMessage and return an error
func (repo SlackMockRepository) EphemeralFileMessage(channel string, user string, message string, title string) error {
	SlackCalls++
	return nil
}

// InitSlackMock returns a SlackMockRepository interface
func InitSlackMock() appcontext.Component {

	return SlackMockRepository{}
}
