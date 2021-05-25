package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/domain"
	"github.com/betorvs/sensubot/utils"
	v2 "github.com/sensu/sensu-go/api/core/v2"
)

// requestSensu func receives a string message, userid and channelid and process it using sensu api
func requestSensu(message string) string {
	if config.Values.SensuAPI == "Absent" || config.Values.SensuAPIToken == "Absent" {
		return "SensuBot Not Configured to access Sensu API"
	}
	sensu := domain.GetSensuRepository()
	command := parseRequest(message)
	var s string
	switch command["verb"] {
	case "get":
		// if config.DebugSensuRequests == "true" {
		// 	log.Printf("[INFO] requestSensu verb get: %s", command["resource"])
		// }
		sensuURL := sensuURLGenerator(command["resource"], command["namespace"], command["name"])
		data, err := sensu.SensuGet(sensuURL, config.Values.SensuAPIToken)
		if err != nil {
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}

		switch command["resource"] {
		case "assets":
			result := new([]v2.Asset)
			// result := new([]domain.SensuAsset)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				if n.Labels != nil {
					s += fmt.Sprintf(" Labels: %s\n", n.Labels)
				}
				if n.Annotations != nil {
					s += fmt.Sprintf(" Annotations: %s\n", n.Annotations)
				}
				s += fmt.Sprintf(" URL: %s\n", n.URL)
				s += fmt.Sprintf(" SHA512: %s\n", n.Sha512)
				// add more info
			}
		case "health":
			result := new(v2.HealthResponse)
			// result := new(domain.Health)
			_ = json.Unmarshal(data, &result)
			for _, v := range result.ClusterHealth {
				s += fmt.Sprintf("Name: %s, Health: %v\n", v.Name, v.Healthy)
			}
			if result.Alarms != nil {
				s += fmt.Sprintf("Alarms: %s\n", result.Alarms)
			}
		case "checks":
			result := new([]v2.Check)
			// result := new([]domain.SensuChecks)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				if n.Labels != nil {
					s += fmt.Sprintf(" Labels: %s\n", n.Labels)
				}
				if n.Annotations != nil {
					s += fmt.Sprintf(" Annotations: %s\n", n.Annotations)
				}
				s += fmt.Sprintf(" Command: %s, Handlers: %s\n", n.Command, n.Handlers)
				s += fmt.Sprintf(" Assets: %s, Subscriptions: %s\n", n.RuntimeAssets, n.Subscriptions)
				// add more info
			}

		case "hooks":
			result := new([]v2.Hook)
			// result := new([]domain.SensuHooks)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				if n.Labels != nil {
					s += fmt.Sprintf("Labels: %s\n", n.Labels)
				}
				if n.Annotations != nil {
					s += fmt.Sprintf("Annotations: %s\n", n.Annotations)
				}
				// add more info
				s += fmt.Sprintf(" Command: %s\n", n.Command)
			}

		case "entities":
			result := new([]v2.Entity)
			// result := new([]domain.SensuEntity)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				if n.Labels != nil {
					s += fmt.Sprintf("Labels: %s\n", n.Labels)
				}
				if n.Annotations != nil {
					s += fmt.Sprintf("Annotations: %s\n", n.Annotations)
				}
				// add more info
				s += fmt.Sprintf(" OS: %s %s, Entity Class: %s\n", n.System.OS, n.System.Platform, n.EntityClass)
				s += fmt.Sprintf(" Subscriptions: %s\n", n.Subscriptions)
			}

		case "handlers":
			result := new([]v2.Handler)
			// result := new([]domain.SensuHandlers)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				if n.Labels != nil {
					s += fmt.Sprintf("Labels: %s\n", n.Labels)
				}
				if n.Annotations != nil {
					s += fmt.Sprintf("Annotations: %s\n", n.Annotations)
				}
				// add more info
				s += fmt.Sprintf(" Command: %s, Handlers: %s\n", n.Command, n.Handlers)
				s += fmt.Sprintf(" Assets: %s, Filters: %s\n", n.RuntimeAssets, n.Filters)
			}

		case "filters":
			result := new([]v2.EventFilter)
			// result := new([]domain.SensuFilter)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				if n.Labels != nil {
					s += fmt.Sprintf("Labels: %s\n", n.Labels)
				}
				if n.Annotations != nil {
					s += fmt.Sprintf("Annotations: %s\n", n.Annotations)
				}
				// add more info
				s += fmt.Sprintf(" Action: %s\n", n.Action)
				s += fmt.Sprintf(" Assets: %s\n", n.RuntimeAssets)
				s += fmt.Sprintf(" Expression: %s\n", n.Expressions)
			}

		case "mutators":
			result := new([]v2.Mutator)
			// result := new([]domain.SensuMutator)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				if n.Labels != nil {
					s += fmt.Sprintf("Labels: %s\n", n.Labels)
				}
				if n.Annotations != nil {
					s += fmt.Sprintf("Annotations: %s\n", n.Annotations)
				}
				// add more info
				s += fmt.Sprintf(" Command: %s\n", n.Command)
				s += fmt.Sprintf(" Assets: %s\n", n.RuntimeAssets)
			}

		case "namespaces":
			result := new([]v2.Namespace)
			// result := new([]domain.SensuNamespaces)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s \n", n.Name)
			}

		case "silenced":
			result := new([]v2.Silenced)
			// result := new([]domain.SensuSilenced)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				if n.Labels != nil {
					s += fmt.Sprintf("Labels: %s\n", n.Labels)
				}
				if n.Annotations != nil {
					s += fmt.Sprintf("Annotations: %s\n", n.Annotations)
				}
				// add more info
				s += fmt.Sprintf(" Creator: %s\n", n.Creator)
				s += fmt.Sprintf(" Subscription: %s\n", n.Subscription)
				s += fmt.Sprintf(" Expire: %d, Expire on Resolve: %v\n", n.Expire, n.ExpireOnResolve)
			}

		case "events":
			result := new([]v2.Event)
			// result := new([]domain.SensuEvents)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Check Name: %s, Namespace: %s\n", n.Check.Name, n.Check.Namespace)
				if n.Check.Labels != nil {
					s += fmt.Sprintf("Labels: %s\n", n.Check.Labels)
				}
				if n.Check.Annotations != nil {
					s += fmt.Sprintf("Annotations: %s\n", n.Check.Annotations)
				}
				// add more info
				s += fmt.Sprintf(" Entity Class: %s\n", n.Entity.EntityClass)
				if n.Check.ProxyEntityName != "" {
					s += fmt.Sprintf(" Proxy Entity Name: %s\n", n.Check.ProxyEntityName)
				}
				s += fmt.Sprintf(" Command: %s\n", n.Check.Command)
				s += fmt.Sprintf(" Assets: %s\n", n.Check.RuntimeAssets)
				s += fmt.Sprintf(" Subscriptions: %s\n", n.Check.Subscriptions)
				s += fmt.Sprintf(" State: %s\n", n.Check.State)
				s += fmt.Sprintf(" Output: %s\n", n.Check.Output)
			}
			// end switch resource
		}

	case "execute":
		sensuURL := fmt.Sprintf("%s/execute", sensuURLGenerator(command["resource"], command["namespace"], command["name"]))
		entity := fmt.Sprintf("entity:%s", command["entity"])
		payload := domain.SensuExecute{
			Check:         command["name"],
			Subscriptions: []string{entity},
		}
		bodymarshal, err := json.Marshal(&payload)
		if err != nil {
			log.Printf("[ERROR] requestSensu read bodyMarshal %s", err)
		}
		_, err = sensu.SensuPost(sensuURL, config.Values.SensuAPIToken, bodymarshal)
		if err != nil {
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}
		s += fmt.Sprintf("Execute Success %s on %s", entity, command["namespace"])

	case "silence":
		sensuURL := sensuURLGenerator("silenced", command["namespace"], "all")
		silencename := fmt.Sprintf("%s:%s", command["entity"], command["name"])

		payload := v2.Silenced{
			ObjectMeta: v2.ObjectMeta{
				Name:      silencename,
				Namespace: command["namespace"],
			},
			Creator:      "sensubot",
			Subscription: command["entity"],
		}
		bodymarshal, err := json.Marshal(&payload)
		if err != nil {
			log.Printf("[ERROR] requestSensu bodyMarshal %s", err)
		}
		_, err = sensu.SensuPost(sensuURL, config.Values.SensuAPIToken, bodymarshal)
		if err != nil {
			log.Printf("[ERROR] requestSensu using sensuPost %s", err)
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output

		}
		s += fmt.Sprintf("Silenced %s Created on %s", command["name"], command["namespace"])
	case "check_bot", "checkbot":
		sensuScheme := config.Values.SensuAPIScheme
		if config.Values.CACertificate != "Absent" || config.Values.SensuAPIScheme != "http" {
			sensuScheme = "https"
		}
		sensuURL := fmt.Sprintf(sensuScheme, config.Values.SensuAPI)
		if !sensu.SensuHealth(sensuURL) {
			return "Sensu API is NOT reachable"
		}
		return "Sensu API is OK"
	}
	// if config.DebugSensuRequests == "true" {
	// 	log.Printf("[INFO] requestSensu Sensu API Response %s", s)
	// }
	return s
}

//sensuURLGenerator func create final sensu url
func sensuURLGenerator(resource string, namespace string, name string) string {
	sensuScheme := config.Values.SensuAPIScheme
	if config.Values.CACertificate != "Absent" || config.Values.SensuAPIScheme != "http" {
		sensuScheme = "https"
	}
	basicURI := "api/core/v2/namespaces"
	if resource == "health" {
		basicURI = string("health")
	} else if resource == "silenced" {
		basicURI = fmt.Sprintf("api/core/v2/namespaces/%s/silenced", namespace)
	} else if resource != "namespaces" {
		basicURI = fmt.Sprintf("api/core/v2/namespaces/%s/%s", namespace, resource)
	} else if name != "all" {
		basicURI = fmt.Sprintf("api/core/v2/namespaces/%s/%s/%s", namespace, resource, name)
	}

	sensuBase := fmt.Sprintf("%s://%s/%s", sensuScheme, config.Values.SensuAPI, basicURI)
	return sensuBase
}

// parseRequest func check if sensubot can understand what is requested
func parseRequest(initialText string) map[string]string {
	space := regexp.MustCompile(`\s+`)
	secondaryText := space.ReplaceAllString(initialText, " ")
	data := strings.Split(strings.ToLower(secondaryText), " ")
	// verb specialWord|resource resouce|name specialWord namespace_name specialWord entity_name
	requestMap := make(map[string]string)
	requestMap["name"] = "all"
	requestMap["namespace"] = "default"
	var specialWords = []string{"on", "to", "in", "at"}
	var verb = []string{"get", "execute", "silence"}
	var resources = []string{"asset", "assets", "check", "checks", "entity", "entities", "handler", "handlers", "hook", "hooks", "filter", "filters", "mutator", "mutators", "namespace", "namespaces", "silence", "silences", "silenced", "health", "event", "events"}
	for k, v := range data {
		switch k {
		case 0:
			if utils.StringInSlice(v, verb) {
				requestMap["verb"] = v
			}

		case 1, 2:
			if utils.StringInSlice(v, verb) {
				requestMap["verb"] = v
			}
			if !utils.StringInSlice(data[k], specialWords) && !utils.StringInSlice(data[k], resources) {
				requestMap["name"] = data[k]
			}
			switch v {
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

			case "event", "events":
				requestMap["resource"] = "events"

			case "health":
				requestMap["resource"] = "health"

			}
		case 3, 4, 5:
			if !utils.StringInSlice(v, specialWords) && !utils.StringInSlice(v, resources) {
				requestMap["namespace"] = v
			}
		case 6:
			if !utils.StringInSlice(data[k], specialWords) && !utils.StringInSlice(data[k], resources) {
				requestMap["entity"] = data[k]
			}
		}

	}

	return requestMap
}
