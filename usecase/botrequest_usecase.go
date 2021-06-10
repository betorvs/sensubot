package usecase

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/utils"
)

// requestToBot func receives a string message, userid and channelid and process it using ToBot api
func requestToBot(message, role, displayName string) string {
	logLocal := config.GetLogger()

	logLocal.Info("Role ", role, " Display Name ", displayName, " and requested ", message)
	command := parseRequest(message)
	logLocal.Debug(command)

	if len(config.Values.BlockedVerbs) != 0 && utils.StringInSlice(command["verb"], config.Values.BlockedVerbs) && role != "admin" {
		logLocal.Debug(config.Values.BlockedVerbs)
		return fmt.Sprintf("Verb %s not allowed", command["verb"])
	}
	if len(config.Values.BlockedResources) != 0 && utils.StringInSlice(command["resource"], config.Values.BlockedResources) && role != "admin" {
		logLocal.Debug(config.Values.BlockedResources)
		return fmt.Sprintf("Resource %s not allowed", command["resource"])
	}
	switch command["integration"] {
	case "sensu":
		logLocal.Debug("using sensu integraton")
		return requestSensu(command, role, displayName)

	case "opsgenie":
		logLocal.Debug("using opsgenie integraton")
		return requestOpsgenie(command, role, displayName)

	case "alertmanager":
		logLocal.Debug("using alertmanager integraton")
		return requestAlertManager(command, role, displayName)
	}
	return fmt.Sprintf("Integration not found %s", command["integration"])
}

// parseRequest func check if sensubot can understand what is requested
func parseRequest(initialText string) map[string]string {
	space := regexp.MustCompile(`\s+`)
	secondaryText := space.ReplaceAllString(strings.TrimLeft(initialText, " "), " ")
	data := strings.Split(secondaryText, " ")
	// verb specialWord|resource resouce|name specialWord namespace_name specialWord entity_name
	requestMap := make(map[string]string)
	requestMap["name"] = "all"
	requestMap["namespace"] = "default"
	var specialWords = []string{"on", "to", "in", "at"}
	var verb = []string{"get", "execute", "silence", "list", "delete", "resolve", "bulkresolve", "bulkdelete"}
	var resources = []string{"asset", "assets", "check", "checks", "entity", "entities", "handler", "handlers", "hook", "hooks", "filter", "filters", "mutator", "mutators", "namespace", "namespaces", "silence", "silences", "silenced", "health", "event", "events", "alerts"}
	namespaceKnown := false
	entityKnown := false
	for k, v := range data {
		value := strings.ToLower(v)
		fmt.Println(k, v)
		switch {
		case k == 0:
			if utils.StringInSlice(value, verb) {
				requestMap["verb"] = value
			}

		case k == 1:
			// test resource to check if it is a integration:
			// examples: opsgenie.alerts or alertmanager.alerts
			integration, resource, valid := utils.ParseKeyValue(value, ".")
			// if is using . to split resource
			if valid {
				// add a new value, like opsgenie, alertmanager
				requestMap["integration"] = integration
				// if this integration name is default
				if integration == "default" {
					// that means it should be configured to use the default integration configured from ENV
					requestMap["integration"] = config.Values.DefaultIntegrationName
				}
			} else {
				resource = value
				// use default integration
				// usually sensu api
				requestMap["integration"] = config.Values.DefaultIntegrationName
			}
			switch resource {
			case "asset", "assets":
				requestMap["resource"] = "assets"

			case "check", "checks":
				requestMap["resource"] = "checks"

			case "hook", "hooks":
				requestMap["resource"] = "hooks"

			case "entity", "entities":
				requestMap["resource"] = "entities"

			case "handler", "handlers":
				requestMap["resource"] = "handlers"

			case "filter", "filters":
				requestMap["resource"] = "filters"

			case "mutator", "mutators":
				requestMap["resource"] = "mutators"

			case "namespace", "namespaces":
				requestMap["resource"] = "namespaces"

			case "silence", "silences", "silenced":
				requestMap["resource"] = "silenced"
				if requestMap["integration"] == "alertmanager" {
					requestMap["resource"] = "silences"
				}

			case "event", "events":
				requestMap["resource"] = "events"

			case "health":
				requestMap["resource"] = "health"

			case "checkbot":
				requestMap["resource"] = "checkbot"

			case "alerts":
				requestMap["resource"] = "alerts"

			}
		case k == 2:
			if !utils.StringInSlice(value, specialWords) && !utils.StringInSlice(value, resources) && !utils.StringInSlice(value, verb) {
				if strings.Contains(data[k], "=") {
					requestMap["labels"] = data[k]
				} else if !strings.Contains(value, ".") {
					requestMap["name"] = data[k]
				}
			}

		case k >= 3:
			if !strings.Contains(v, "=") {
				if !utils.StringInSlice(value, specialWords) && !utils.StringInSlice(value, resources) && !namespaceKnown {
					requestMap["namespace"] = v
					namespaceKnown = true
					continue
				}
				if !utils.StringInSlice(data[k], specialWords) && !utils.StringInSlice(data[k], resources) && !entityKnown {
					requestMap["entity"] = data[k]
					entityKnown = true
					continue
				}
			} else {
				requestMap["labels"] += fmt.Sprintf(",%s", data[k])
			}

		}

	}
	// fmt.Printf("%#v", requestMap)
	return requestMap
}
