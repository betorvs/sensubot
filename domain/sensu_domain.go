package domain

import "github.com/betorvs/sensubot/appcontext"

// SensuExecute struct
type SensuExecute struct {
	Check         string   `json:"check"`
	Subscriptions []string `json:"subscriptions"`
}

// SensuRepository interface
type SensuRepository interface {
	appcontext.Component
	// SensuPost func return []byte and error from a POST using sensu api token
	SensuPost(token string, sensuurl string, body []byte) ([]byte, error)
	// SensuGet func return []byte and error from a requested URL using a sensu api token
	SensuGet(sensuurl string, token string) ([]byte, error)
	// SensuHealth test if sensu api is health
	SensuHealth(sensuurl string) bool
}

// GetSensuRepository func return SensuRepository interface
func GetSensuRepository() SensuRepository {
	return appcontext.Current.Get(appcontext.SensuRepository).(SensuRepository)
}
