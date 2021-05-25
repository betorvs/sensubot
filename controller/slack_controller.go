package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/usecase"

	"github.com/labstack/echo/v4"
	"github.com/slack-go/slack"
)

const (
	// Disabled constante is used to check if a config variable is not configured
	disabled string = "disabled"
)

// SlackEvents func receive an post from slack slash integration
func SlackEvents(c echo.Context) error {
	if config.Values.SlackToken == disabled || config.Values.SlackSigningSecret == disabled {
		return c.JSON(http.StatusNotImplemented, nil)
	}
	var bodyBytes []byte
	if c.Request().Body == nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)
	// Copy Forms to Struct
	data := new(slack.SlashCommand)
	data.Token = c.FormValue("token")
	data.TeamID = c.FormValue("team_id")
	data.TeamDomain = c.FormValue("team_domain")
	data.EnterpriseID = c.FormValue("enterprise_id")
	data.EnterpriseName = c.FormValue("enterprise_name")
	data.ChannelID = c.FormValue("channel_id")
	data.ChannelName = c.FormValue("channel_name")
	data.UserID = c.FormValue("user_id")
	data.UserName = c.FormValue("user_name")
	data.Command = c.FormValue("command")
	data.Text = c.FormValue("text")
	data.ResponseURL = c.FormValue("response_url")
	data.TriggerID = c.FormValue("trigger_id")
	// Headers
	slackRequestTimestamp := c.Request().Header.Get("X-Slack-Request-Timestamp")
	slackSignature := c.Request().Header.Get("X-Slack-Signature")

	basestring := fmt.Sprintf("v0:%s:%s", slackRequestTimestamp, bodyString)
	verifier := usecase.ValidateBot(slackSignature, basestring, config.Values.SlackSigningSecret)
	if !verifier {
		return c.JSON(http.StatusForbidden, nil)
	}
	if data.ChannelID != config.Values.SlackChannel {
		return c.JSON(http.StatusForbidden, nil)
	}
	if data.Command != config.Values.SlackSlashCommand {
		return c.JSON(http.StatusForbidden, nil)
	}
	res, err := usecase.ParseSlashCommand(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusAccepted, res)
}
