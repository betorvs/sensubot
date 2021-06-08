package telegramclient

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
)

// TelegramRepository struct
type TelegramRepository struct {
	Client *http.Client
}

// SendTelegramMessage takes a chatID and sends answer to them
func (repo TelegramRepository) SendTelegramMessage(reqBody []byte) error {
	logLocal := config.GetLogger()
	logLocal.Debug("starting SendTelegramMessage")
	// Send a post request with your token
	telegramURL := fmt.Sprintf("%s%s/sendMessage", config.Values.TelegramURL, config.Values.TelegramToken)
	req, err := http.NewRequest(http.MethodPost, telegramURL, bytes.NewBuffer(reqBody))
	if err != nil {
		logLocal.Error(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := repo.Client.Do(req)
	if err != nil {
		logLocal.Error(err)
		return err
	}
	if res.StatusCode != http.StatusOK {
		logLocal.Error(err)
		return errors.New("unexpected status" + res.Status)
	}
	defer res.Body.Close()
	return nil
}

func New() appcontext.Component {
	client := http.Client{
		Timeout: time.Second * time.Duration(config.Values.BotTimeout),
	}
	return &TelegramRepository{Client: &client}
}

func init() {
	if config.Values.TestRun {
		return
	}

	if config.Values.TelegramToken != "disabled" {
		appcontext.Current.Add(appcontext.TelegramRepository, New)
		logLocal := config.GetLogger()
		logLocal.Info("Connected to Telegram")
		logLocal.Debug("Telegram Admin list ", config.Values.TelegramAdminIDList)
	}

}
