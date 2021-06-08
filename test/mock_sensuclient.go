package test

import "github.com/betorvs/sensubot/appcontext"

var (
	// SensuCalls int
	SensuCalls int
	// SensuPostCalls int
	SensuPostCalls int
	// SensuHealthCalls int
	SensuHealthCalls int
	// SensuDeleteCalls int
	SensuDeleteCalls int
)

// SensuMockRepository struct mock
type SensuMockRepository struct {
}

// SensuGet func return []byte and error from a requested URL using a sensu api token
func (repo SensuMockRepository) SensuGet(sensuurl string) ([]byte, error) {
	SensuCalls++
	return []byte{}, nil
}

// SensuPost func return []byte and error from a POST using sensu api token
func (repo SensuMockRepository) SensuPost(sensuurl string, method string, body []byte) ([]byte, error) {
	SensuPostCalls++
	return []byte{}, nil
}

// SensuHealth func
func (repo SensuMockRepository) SensuHealth(sensuurl string) bool {
	SensuHealthCalls++
	return true
}

func (repo SensuMockRepository) SensuDelete(sensuURL string) error {
	SensuDeleteCalls++
	return nil
}

// InitSensuMock returns a SensuMockRepository interface
func InitSensuMock() appcontext.Component {

	return SensuMockRepository{}
}
