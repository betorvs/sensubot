package domain

import (
	"github.com/betorvs/sensubot/appcontext"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/incident"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

// OpsgenieRepository interface
type OpsgenieRepository interface {
	appcontext.Component
	// GetAlert by tinyID string and return *alert.GetAlertResult, error
	GetAlert(tinyID string) (*alert.GetAlertResult, error)
	// ListAlert by query string and return  (*alert.ListAlertResult, error
	ListAlert(query string) (*alert.ListAlertResult, error)
	// ListIncident by query string and return *incident.ListResult, error
	ListIncident(query string) (*incident.ListResult, error)
	// GetOnCall(team string) (*schedule.GetOnCallsResult, error)
	GetOnCall(team string) (*schedule.GetOnCallsResult, error)
	// ListSchedules() (*schedule.ListResult, error)
	ListSchedules() (*schedule.ListResult, error)
}

// GetOpsgenieRepository func return OpsgenieRepository interface
func GetOpsgenieRepository() OpsgenieRepository {
	return appcontext.Current.Get(appcontext.OpsgenieRepository).(OpsgenieRepository)
}
