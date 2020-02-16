package telegramclient

import (
	"bytes"
	"errors"
	"fmt"
	"log"
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

	// Send a post request with your token
	telegramURL := fmt.Sprintf("%s%s/sendMessage", config.TelegramURL, config.TelegramToken)
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

func init() {
	if config.GetEnv("TESTRUN", "false") == "true" {
		return
	}
	client := http.Client{
		Timeout: time.Second * config.BotTimeout,
	}
	appcontext.Current.Add(appcontext.TelegramRepository, TelegramRepository{Client: &client})
	if appcontext.Current.Count() != 0 {
		log.Println("[INFO] Telegram Repository initiated")
	}
}
