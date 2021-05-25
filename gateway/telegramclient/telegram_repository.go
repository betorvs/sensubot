package telegramclient

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
)

// TelegramRepository struct
type TelegramRepository struct {
	Client *http.Client
}

// SendTelegramMessage takes a chatID and sends answer to them
func (repo TelegramRepository) SendTelegramMessage(reqBody []byte) error {

	// Send a post request with your token
	telegramURL := fmt.Sprintf("%s%s/sendMessage", config.Values.TelegramURL, config.Values.TelegramToken)
	req, err := http.NewRequest("POST", telegramURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := repo.Client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}
	defer res.Body.Close()
	return nil
}

func New() appcontext.Component {
	client := http.Client{
		Timeout: config.Values.BotTimeout,
	}
	return &TelegramRepository{Client: &client}
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.TelegramRepository, New)
	if config.Values.TelegramToken != "disabled" {
		logLocal := config.GetLogger()
		logLocal.Info("Connected to Telegram")
	}

}
