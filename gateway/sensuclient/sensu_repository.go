package sensuclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
)

// SensuRepository struct
type SensuRepository struct {
	Client *http.Client
}

// SensuGet func return []byte and error from a requested URL using a sensu api token
func (repo SensuRepository) SensuGet(sensuurl string, token string) ([]byte, error) {
	if config.DebugSensuRequests == "true" {
		log.Printf("[INFO] sensuGet started: %s", sensuurl)
	}
	req, err := http.NewRequest("GET", sensuurl, nil)
	if err != nil {
		return []byte{}, err
	}
	var bearer = "Key " + token
	req.Header.Add("Authorization", bearer)
	resp, err := repo.Client.Do(req)
	if err != nil {
		log.Println("[ERROR] sensuGet create Client")
		return []byte{}, err
	}
	if config.DebugSensuRequests == "true" {
		log.Printf("[INFO] sensuGet response: %s", resp.Status)
	}
	if resp.StatusCode != 200 {
		log.Printf("[ERROR] sensuGet response %s", resp.Status)
		err = fmt.Errorf("%s", resp.Status)
		return []byte{}, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] sensuGet read body")
		return []byte{}, err
	}
	defer resp.Body.Close()
	return bodyText, nil
}

// SensuPost func return []byte and error from a POST using sensu api token
func (repo SensuRepository) SensuPost(sensuurl string, token string, body []byte) ([]byte, error) {
	if config.DebugSensuRequests == "true" {
		log.Printf("[INFO]: sensuPost started: %s", sensuurl)
	}
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest("POST", sensuurl, bytes.NewBuffer(body))
		if err != nil {
			log.Println("[ERROR] sensuPost newRequest")
			return []byte{}, err
		}
	} else {
		req, err = http.NewRequest("POST", sensuurl, nil)
		if err != nil {
			log.Println("[ERROR] sensuPost newRequest")
			return []byte{}, err
		}
	}

	var bearer = fmt.Sprintf("Key %s", token)
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := repo.Client.Do(req)
	if err != nil {
		log.Println("[ERROR] sensuPost create Client")
		return []byte{}, err
	}
	if config.DebugSensuRequests == "true" {
		log.Printf("[INFO] sensuPost response: %s", resp.Status)
	}
	if resp.StatusCode > 204 {
		log.Printf("[ERROR] sensuPost response: %s", resp.Status)
		err = fmt.Errorf("%s", resp.Status)
		return []byte{}, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] sensuPost read body")
		return []byte{}, err
	}

	defer resp.Body.Close()
	return bodyText, nil
}

// SensuHealth func
func (repo SensuRepository) SensuHealth(sensuurl string) bool {
	sensuURL := fmt.Sprintf("%s/health", sensuurl)
	req, err := http.NewRequest("GET", sensuURL, nil)
	if err != nil {
		return false
	}
	resp, err := repo.Client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true
}

func init() {
	if config.GetEnv("TESTRUN", "false") == "true" {
		return
	}
	client := http.Client{
		Timeout: time.Second * config.BotTimeout,
	}
	if config.CACertificate != "Absent" {
		caCert, err := ioutil.ReadFile(config.CACertificate)
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		client = http.Client{
			Timeout: time.Second * config.BotTimeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
		}
	}
	appcontext.Current.Add(appcontext.SensuRepository, SensuRepository{Client: &client})
	if appcontext.Current.Count() != 0 {
		log.Println("[INFO] Sensu Repository initiated")
	}

}
