package test

import "github.com/betorvs/sensubot/appcontext"

var (
	// TelegramCalls int
	TelegramCalls int
)

// TelegramMockRepository struct mock
type TelegramMockRepository struct {
}

// SendTelegramMessage takes a chatID and sends answer to them
func (repo TelegramMockRepository) SendTelegramMessage(reqBody []byte) error {
	TelegramCalls++
	return nil
}

// InitTelegramMock returns a TelegramMockRepository interface
func InitTelegramMock() appcontext.Component {

	return TelegramMockRepository{}
}
