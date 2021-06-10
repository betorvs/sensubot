package alertmanagerclient

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
)

// AlertManagerRepository struct
type AlertManagerRepository struct {
	Client *http.Client
}

// GetAlertManager get http alerts from AM
func (repo AlertManagerRepository) GetAlertManager(amURL string) (result []byte, err error) {
	logLocal := config.GetLogger()
	req, err := http.NewRequest(http.MethodGet, amURL, nil)
	if err != nil {
		logLocal.Error("GET ", err)
		return nil, err
	}
	if config.Values.AlertManagerUsername != "Absent" && config.Values.AlertManagerPassword != "Absent" {
		req.SetBasicAuth(config.Values.AlertManagerUsername, config.Values.AlertManagerPassword)
	}
	resp, err := repo.Client.Do(req)
	if err != nil {
		logLocal.Error("client ", err)
		return nil, err
	}
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logLocal.Error("ReadAll ", err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

func New() appcontext.Component {
	client := http.Client{
		Timeout: time.Second * time.Duration(config.Values.BotTimeout),
	}
	return &AlertManagerRepository{Client: &client}
}

func init() {
	if config.Values.TestRun {
		return
	}

	if len(config.Values.AlertManagerEndpoints) != 0 {
		appcontext.Current.Add(appcontext.AlertManagerRepository, New)
		logLocal := config.GetLogger()
		logLocal.Info("Connected to Alert Manager Endpoints")
	}

}
