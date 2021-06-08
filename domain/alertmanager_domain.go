package domain

import "github.com/betorvs/sensubot/appcontext"

// AlertManagerRepository interface
type AlertManagerRepository interface {
	appcontext.Component
	// GetAlertManager receives an string URL and return []byte and error
	GetAlertManager(amURL string) (result []byte, err error)
}

// GetAlertManagerRepository func return AlertManagerRepository interface
func GetAlertManagerRepository() AlertManagerRepository {
	return appcontext.Current.Get(appcontext.AlertManagerRepository).(AlertManagerRepository)
}
