package controller

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/usecase"
	oidc "github.com/coreos/go-oidc"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/chat/v1"
)

const (
	jwtURL     string = "https://www.googleapis.com/service_accounts/v1/jwk/"
	chatIssuer string = "chat@system.gserviceaccount.com"
)

// GChatEvents func receive an post from GChat slash integration
func GChatEvents(c echo.Context) error {
	// "Authorization"
	bearerToken := c.Request().Header.Get("Authorization")
	// fmt.Println(bearerToken)
	token := strings.Split(bearerToken, " ")
	if len(token) != 2 || !verifyJWT(config.Values.GoogleChatProjectID, token[1]) {
		// fmt.Println(len(token), verifyJWT(config.Values.GoogleChatProjectID, token[1]))
		return c.JSON(http.StatusForbidden, nil)
	}
	event := new(chat.DeprecatedEvent)
	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	switch event.Type {
	case "ADDED_TO_SPACE":
		if event.Space.Type != "ROOM" {
			break
		}
		text := new(chat.Message)
		text.Text = "thanks for adding me."
		return c.JSON(http.StatusAccepted, text)
	case "MESSAGE":
		text := new(chat.Message)

		res, err := usecase.ParseGoogleChatCommand(event)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		text.Text = fmt.Sprintf("you requested %s %s", event.Message.ArgumentText, res)
		return c.JSON(http.StatusAccepted, text)
	}
	return nil
}

func verifyJWT(audience, token string) bool {
	logLocal := config.GetLogger()
	logLocal.Debug("starting verifying")
	ctx := context.Background()
	keySet := oidc.NewRemoteKeySet(ctx, jwtURL+chatIssuer)

	configLocal := &oidc.Config{
		SkipClientIDCheck: false,
		ClientID:          audience,
	}
	newVerifier := oidc.NewVerifier(chatIssuer, keySet, configLocal)
	test, err := newVerifier.Verify(ctx, token)
	if err != nil {
		logLocal.Debug("Audience dont match")
		return false
	}
	if len(test.Audience) == 1 {
		logLocal.Debug("Audience matches")
	}

	return true
}
