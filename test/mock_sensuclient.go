package test

import "github.com/betorvs/sensubot/appcontext"

var (
	// SensuCalls int
	SensuCalls int
)

// SensuMockRepository struct mock
type SensuMockRepository struct {
}

// SensuGet func return []byte and error from a requested URL using a sensu api token
func (repo SensuMockRepository) SensuGet(sensuurl string, token string) ([]byte, error) {
	SensuCalls++
	return []byte{}, nil
}

// SensuPost func return []byte and error from a POST using sensu api token
func (repo SensuMockRepository) SensuPost(sensuurl string, token string, body []byte) ([]byte, error) {
	SensuCalls++
	return []byte{}, nil
}

// SensuHealth func
func (repo SensuMockRepository) SensuHealth(sensuurl string) bool {
	SensuCalls++
	return true
}

// SendSensuMessage takes a chatID and sends answer to them
// func (repo TelegramMockRepository) SendTelegramMessage(reqBody []byte) error {
// 	TelegramCalls++
// 	return nil
// }

// InitSensuMock returns a SensuMockRepository interface
func InitSensuMock() appcontext.Component {

	return SensuMockRepository{}
}
