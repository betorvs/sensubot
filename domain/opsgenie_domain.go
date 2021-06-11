package domain

import (
	"github.com/betorvs/sensubot/appcontext"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

// OpsgenieRepository interface
type OpsgenieRepository interface {
	appcontext.Component
	// GetAlert by tinyID string and return *alert.GetAlertResult, error
	GetAlert(tinyID string) (*alert.GetAlertResult, error)
	// ListAlert by query string and return  (*alert.ListAlertResult, error
	ListAlert(query string) (*alert.ListAlertResult, error)
}

// GetOpsgenieRepository func return OpsgenieRepository interface
func GetOpsgenieRepository() OpsgenieRepository {
	return appcontext.Current.Get(appcontext.OpsgenieRepository).(OpsgenieRepository)
}
