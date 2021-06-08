package sensuclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
)

// SensuRepository struct
type SensuRepository struct {
	Client *http.Client
}

// Auth represents the authentication info
type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

var currentToken = Auth{}

// SensuGet func return []byte and error from a requested URL using a sensu api token
func (repo SensuRepository) SensuGet(sensuurl string) ([]byte, error) {
	logLocal := config.GetLogger()
	logLocal.Debugf("sensuGet started %s", sensuurl)
	req, err := http.NewRequest(http.MethodGet, sensuurl, nil)
	if err != nil {
		return []byte{}, err
	}
	if config.Values.SensuAPIToken == "Absent" {
		auth, autherr := authenticate(repo.Client)
		if autherr != nil {
			return []byte{}, autherr
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth.AccessToken))
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Key %s", config.Values.SensuAPIToken))
	}
	resp, err := repo.Client.Do(req)
	if err != nil {
		logLocal.Error("sensuGet create Client")
		return []byte{}, err
	}
	logLocal.Debugf("sensuGet response %s", resp.Status)
	if resp.StatusCode != 200 {
		logLocal.Errorf("sensuGet response %s", resp.Status)
		err = fmt.Errorf("%s", resp.Status)
		return []byte{}, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logLocal.Error("sensuGet read body ", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	return bodyText, nil
}

// SensuPost func return []byte and error from a POST using sensu api token
func (repo SensuRepository) SensuPost(sensuurl string, method string, body []byte) ([]byte, error) {
	logLocal := config.GetLogger()
	logLocal.Debugf("sensuPost started %s", sensuurl)
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest(method, sensuurl, bytes.NewBuffer(body))
		if err != nil {
			logLocal.Error("sensuPost newRequest with body")
			return []byte{}, err
		}
	} else {
		req, err = http.NewRequest(method, sensuurl, nil)
		if err != nil {
			logLocal.Error("sensuPost newRequest without body")
			return []byte{}, err
		}
	}

	if config.Values.SensuAPIToken == "Absent" {
		auth, autherr := authenticate(repo.Client)
		if autherr != nil {
			return []byte{}, autherr
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth.AccessToken))
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Key %s", config.Values.SensuAPIToken))
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := repo.Client.Do(req)
	if err != nil {
		logLocal.Error("sensuPost create Client")
		return []byte{}, err
	}
	logLocal.Debugf("sensuPost response %s", resp.Status)
	if resp.StatusCode > 204 {
		logLocal.Errorf("sensuPost response %s", resp.Status)
		err = fmt.Errorf("%s", resp.Status)
		return []byte{}, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logLocal.Error("sensuPost read body")
		return []byte{}, err
	}

	defer resp.Body.Close()
	return bodyText, nil
}

// SensuHealth func
func (repo SensuRepository) SensuHealth(sensuurl string) bool {
	sensuURL := fmt.Sprintf("%s/health", sensuurl)
	req, err := http.NewRequest(http.MethodGet, sensuURL, nil)
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

// SensuDelete func
func (repo SensuRepository) SensuDelete(sensuURL string) error {

	req, err := http.NewRequest(http.MethodDelete, sensuURL, nil)
	if err != nil {
		return fmt.Errorf("error deleting request for %s: %v", sensuURL, err)
	}
	if config.Values.SensuAPIToken == "Absent" {
		auth, autherr := authenticate(repo.Client)
		if autherr != nil {
			return autherr
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth.AccessToken))
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Key %s", config.Values.SensuAPIToken))
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := repo.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing DELETE request for %s: %v", sensuURL, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("Delete check to %s failed with status %v ", sensuURL, resp.Status)
	}

	return nil
}

// authenticate funcion to wotk with api-backend-* flags
func authenticate(client *http.Client) (Auth, error) {
	if currentToken.AccessToken != "" && currentToken.ExpiresAt > time.Now().Unix() {
		return currentToken, nil
	}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s://%s/auth", config.Values.SensuAPIScheme, config.Values.SensuAPI),
		nil,
	)
	if err != nil {
		return currentToken, fmt.Errorf("error generating auth request: %v", err)
	}

	req.SetBasicAuth(config.Values.SensuUser, config.Values.SensuPassword)

	resp, err := client.Do(req)
	if err != nil {
		return currentToken, fmt.Errorf("error executing auth request: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return currentToken, fmt.Errorf("error reading auth response: %v", err)
	}

	if strings.HasPrefix(string(body), "Unauthorized") {
		return currentToken, fmt.Errorf("authorization failed for user %s", config.Values.SensuUser)
	}

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&currentToken)

	if err != nil {
		return currentToken, fmt.Errorf("error decoding auth response: %v\nFirst response: %s", err, body)
	}

	return currentToken, err
}

func New() appcontext.Component {
	client := http.Client{
		Timeout: time.Second * time.Duration(config.Values.BotTimeout),
	}
	if config.Values.CACertificate != "Absent" {
		caCert, err := ioutil.ReadFile(config.Values.CACertificate)
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		client = http.Client{
			Timeout: time.Second * time.Duration(config.Values.BotTimeout),
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
		}
	}
	return &SensuRepository{Client: &client}
}

func init() {
	if config.Values.TestRun {
		return
	}

	if config.Values.SensuAPI != "Absent" || config.Values.SensuAPIToken != "Absent" {
		appcontext.Current.Add(appcontext.SensuRepository, New)
		logLocal := config.GetLogger()
		logLocal.Info("Connected to Sensu API with timeout ", config.Values.BotTimeout)
	}

}
