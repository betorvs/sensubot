package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/domain"
	"github.com/betorvs/sensubot/utils"
	"github.com/sensu/sensu-go/types"
)

// requestSensu func receives a map[string]string command, role and display Name and process it using sensu api
func requestSensu(command map[string]string, role, displayName string) string {
	if !checkSensuAPIConfig() {
		return "SensuBot Not Configured to access Sensu API"
	}
	logLocal := config.GetLogger()

	sensu := domain.GetSensuRepository()
	var s string
	switch command["verb"] {
	case "get":
		var k, v string
		var b bool
		if command["labels"] != "" {
			k, v, b = utils.ParseKeyValue(command["labels"], "=")
		}
		sensuURL := sensuURLGenerator(command["resource"], command["namespace"], command["name"])
		logLocal.Debug(sensuURL)
		data, err := sensu.SensuGet(sensuURL)
		if err != nil {
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}
		switch command["resource"] {
		case "assets":
			result := new([]types.Asset)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
			result := new(types.HealthResponse)
			_ = json.Unmarshal(data, &result)
			for _, v := range result.ClusterHealth {
				s += fmt.Sprintf("Name: %s, Health: %v\n", v.Name, v.Healthy)
			}
			if result.Alarms != nil {
				s += fmt.Sprintf("Alarms: %s\n", result.Alarms)
			}
		case "checks":
			result := new([]types.Check)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
			result := new([]types.Hook)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
			result := new([]types.Entity)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
			result := new([]types.Handler)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
			result := new([]types.EventFilter)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
			result := new([]types.Mutator)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
			result := new([]types.Namespace)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				s += fmt.Sprintf("Name: %s \n", n.Name)
			}

		case "silenced":
			result := new([]types.Silenced)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
			result := new([]types.Event)
			_ = json.Unmarshal(data, &result)
			for _, n := range *result {
				if b {
					if !utils.SearchLabels(k, v, n.Labels) {
						continue
					}
				}
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
		// end get verb

	case "list":
		var k, v, r string
		var b bool
		if command["labels"] != "" {
			k, v, b = utils.ParseKeyValue(command["labels"], "=")
		}
		if !b {
			output := "Sorry. List request should be used with labels"
			return output
		}
		sensuURL := sensuURLGenerator(command["resource"], command["namespace"], command["name"])
		logLocal.Debug(sensuURL)
		data, err := sensu.SensuGet(sensuURL)
		if err != nil {
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}
		switch command["resource"] {
		case "assets":
			result := new([]types.Asset)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				r += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				count++
			}
			s += fmt.Sprintf("Number of Assets found: %v \n", count)
			s += r

		case "checks":
			result := new([]types.Check)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				r += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				count++
			}
			s += fmt.Sprintf("Number of Checks found: %v \n", count)
			s += r

		case "hooks":
			result := new([]types.Hook)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				r += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				count++
			}
			s += fmt.Sprintf("Number of Hooks found: %v \n", count)
			s += r

		case "entities":
			result := new([]types.Entity)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				r += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				count++
			}
			s += fmt.Sprintf("Number of Entities found: %v \n", count)
			s += r

		case "handlers":
			result := new([]types.Handler)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				r += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				count++
			}
			s += fmt.Sprintf("Number of Handlers found: %v \n", count)
			s += r

		case "filters":
			result := new([]types.EventFilter)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				r += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				count++
			}
			s += fmt.Sprintf("Number of Filters found: %v \n", count)
			s += r

		case "mutators":
			result := new([]types.Mutator)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				r += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				count++
			}
			s += fmt.Sprintf("Number of Mutators found: %v \n", count)
			s += r

		case "namespaces":
			result := new([]types.Namespace)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				r += fmt.Sprintf("Name: %s \n", n.Name)
				count++
			}
			s += fmt.Sprintf("Number of Namespaces found: %v \n", count)
			s += r

		case "silenced":
			result := new([]types.Silenced)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				r += fmt.Sprintf("Name: %s, Namespace: %s\n", n.Name, n.Namespace)
				count++
			}
			s += fmt.Sprintf("Number of Silence found: %v \n", count)
			s += r

		case "events":
			result := new([]types.Event)
			_ = json.Unmarshal(data, &result)
			count := 0
			for _, n := range *result {
				var checked bool
				if utils.SearchLabels(k, v, n.Check.Labels) {
					count++
					r += fmt.Sprintf("Check Name: %s, Entity: %s, Namespace: %s\n", n.Check.Name, n.Entity.Name, n.Check.Namespace)
					checked = true
				}
				if utils.SearchLabels(k, v, n.Entity.Labels) && !checked {
					count++
					r += fmt.Sprintf("Check Name: %s, Entity: %s, Namespace: %s\n", n.Check.Name, n.Entity.Name, n.Check.Namespace)
					checked = true
				}
				if utils.SearchLabels(k, v, n.Labels) && !checked {
					count++
					r += fmt.Sprintf("Check Name: %s, Entity: %s, Namespace: %s\n", n.Check.Name, n.Entity.Name, n.Check.Namespace)
					checked = true
				}
			}
			s += fmt.Sprintf("Number of Events found: %v \n", count)
			s += r
			// end switch resource
		}
		// end list verb

	case "execute":
		sensuURL := sensuURLGenerator(command["verb"], command["namespace"], command["name"])
		logLocal.Debug(sensuURL)
		entity := fmt.Sprintf("entity:%s", command["entity"])
		payload := domain.SensuExecute{
			Check:         command["name"],
			Subscriptions: []string{entity},
		}
		bodymarshal, err := json.Marshal(&payload)
		if err != nil {
			logLocal.Errorf("execute verb requestSensu read bodyMarshal %s", err)
		}
		_, err = sensu.SensuPost(sensuURL, "POST", bodymarshal)
		if err != nil {
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}
		s += fmt.Sprintf("Execute Success %s on %s", entity, command["namespace"])

	case "silence":
		silencename := fmt.Sprintf("%s:%s", command["entity"], command["name"])
		sensuURL := sensuURLGenerator("silenced", command["namespace"], silencename)
		logLocal.Debug(sensuURL)

		payload := types.Silenced{
			ObjectMeta: types.ObjectMeta{
				Name:      silencename,
				Namespace: command["namespace"],
				Labels: map[string]string{
					config.Values.AppName: "owner",
				},
			},
			Reason: fmt.Sprintf("Requested by %s using %s", displayName, config.Values.AppName),
		}

		if command["name"] != "all" {
			payload.Check = command["name"]
		} else {
			payload.Subscription = fmt.Sprintf("entity:%s", command["entity"])
		}
		bodymarshal, err := json.Marshal(&payload)
		if err != nil {
			logLocal.Errorf("silence verb requestSensu bodyMarshal %s", err)
		}
		_, err = sensu.SensuPost(sensuURL, "PUT", bodymarshal)
		if err != nil {
			logLocal.Errorf("silence verb requestSensu using sensuPost with PUT %s", err)
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}
		s += fmt.Sprintf("Silenced %s Created on %s", command["name"], command["namespace"])

	case "checkbot":
		sensuScheme := config.Values.SensuAPIScheme
		if config.Values.CACertificate != "Absent" || config.Values.SensuAPIScheme != "http" {
			sensuScheme = "https"
		}
		sensuURL := fmt.Sprintf(sensuScheme, config.Values.SensuAPI)
		if !sensu.SensuHealth(sensuURL) {
			return "Sensu API is NOT reachable"
		}
		return "Sensu API is OK"

	case "delete":
		// clean resources like events and checks
		if command["name"] == "all" {
			output := "Sorry. Request Failed. Resource name missing"
			return output
		}
		if !utils.StringInSlice(command["resource"], config.Values.DeleteAllowedResources) {
			output := fmt.Sprintf("Sorry. Request Failed. Resource requested to be deleted not allowed %s", command["resource"])
			return output
		}
		sensuURL := sensuURLGenerator(command["resource"], command["namespace"], command["name"])
		logLocal.Debug(sensuURL)
		err := sensu.SensuDelete(sensuURL)
		if err != nil {
			output := fmt.Sprintf("Sorry. Delete request Failed: %s", err)
			return output
		}
		s += fmt.Sprintf("Delete %s Successfully %s on %s", command["resource"], command["name"], command["namespace"])

	case "resolve":
		sensuURL := fmt.Sprintf("%s/%s", sensuURLGenerator("events", command["namespace"], command["entity"]), command["name"])
		logLocal.Debug(sensuURL)
		data, err := sensu.SensuGet(sensuURL)
		if err != nil {
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}
		result := new(types.Event)
		_ = json.Unmarshal(data, &result)
		output := fmt.Sprintf("Resolved by %s via %s \n %s", displayName, config.Values.AppName, result.Check.Output)
		result.Check.Output = output
		result.Check.Status = uint32(0)
		localSensuURL := fmt.Sprintf("%s/%s", sensuURLGenerator("events", command["namespace"], result.Entity.Name), result.Check.Name)
		bodymarshal, err := json.Marshal(&result)
		if err != nil {
			logLocal.Errorf("resolve verb requestSensu bodyMarshal %s", err)
		}
		_, err = sensu.SensuPost(localSensuURL, "PUT", bodymarshal)
		if err != nil {
			logLocal.Errorf("resolve verb requestSensu using sensuPost with PUT %s", err)
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}

		s += fmt.Sprintf("Resolved Check Name: %s, Entity: %s, Namespace: %s\n", result.Check.Name, result.Entity.Name, result.Check.Namespace)

	case "bulkResolve", "bulkresolve":
		var k, v string
		var b bool
		if command["labels"] != "" {
			k, v, b = utils.ParseKeyValue(command["labels"], "=")
		}
		if !b {
			output := "Sorry. Resolve request should be used with labels"
			return output
		}
		sensuURL := sensuURLGenerator("events", command["namespace"], "all")
		logLocal.Debug(sensuURL)
		data, err := sensu.SensuGet(sensuURL)
		if err != nil {
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}
		result := new([]types.Event)
		_ = json.Unmarshal(data, &result)
		s += fmt.Sprintf("Total of Events found: %v \n", len(*result))
		count := 0
		for _, n := range *result {
			var checked bool
			if !utils.SearchLabels(k, v, n.Check.Labels) {
				checked = true
			}
			if !utils.SearchLabels(k, v, n.Entity.Labels) && !checked {
				checked = true
			}
			if !utils.SearchLabels(k, v, n.Labels) && !checked {
				checked = true
			}
			if checked {
				continue
			}
			output := fmt.Sprintf("Resolved by %s via %s \n %s", displayName, config.Values.AppName, n.Check.Output)
			n.Check.Output = output
			n.Check.Status = uint32(0)
			localSensuURL := fmt.Sprintf("%s/%s", sensuURLGenerator("events", command["namespace"], n.Entity.Name), n.Check.Name)
			bodymarshal, err := json.Marshal(&n)
			if err != nil {
				logLocal.Errorf("resolve verb requestSensu bodyMarshal %s", err)
			}
			_, err = sensu.SensuPost(localSensuURL, "PUT", bodymarshal)
			if err != nil {
				logLocal.Errorf("resolve verb requestSensu using sensuPost with PUT %s", err)
				output := fmt.Sprintf("Sorry. Request Failed: %s", err)
				return output
			}
			count++

			// s += fmt.Sprintf("Check Name: %s, Entity: %s, Namespace: %s\n", n.Check.Name, n.Entity.Name, n.Check.Namespace)
		}
		s += fmt.Sprintf("Number of Events resolved: %v \n", count)
		if count == len(*result) {
			s += fmt.Sprintf("All Events were resolved: %v \n", len(*result))
		}
	case "bulkDelete", "bulkdelete":
		// very usefull if using integrations like
		// sensu-alertmanager-events
		// sensu-kubernetes-events
		// sensu-dynamic-check-mutator
		// name here will be used as key to search in labels from that resource
		bulkDeleteResources := []string{"events", "checks"}
		if !utils.StringInSlice(command["resource"], bulkDeleteResources) {
			output := "Sorry. bulkDelete request can only be used with events and checks"
			return output
		}
		var k, v string
		var b bool
		if command["labels"] != "" {
			k, v, b = utils.ParseKeyValue(command["labels"], "=")
		}
		if !b {
			output := "Sorry. List request should be used with labels"
			return output
		}
		sensuURL := sensuURLGenerator(command["resource"], command["namespace"], "all")
		logLocal.Debug(sensuURL)
		data, err := sensu.SensuGet(sensuURL)
		if err != nil {
			output := fmt.Sprintf("Sorry. Request Failed: %s", err)
			return output
		}
		switch command["resource"] {
		case "events":
			result := new([]types.Event)
			_ = json.Unmarshal(data, &result)
			s += fmt.Sprintf("Total of %s found: %v", command["resource"], len(*result))
			logLocal.Debug("Total of ", command["resource"], " found: ", len(*result))
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				localSensuURL := fmt.Sprintf("%s/%s", sensuURLGenerator("events", command["namespace"], n.Entity.Name), n.Check.Name)
				err = sensu.SensuDelete(localSensuURL)
				if err != nil {
					logLocal.Errorf("delete verb requestSensu using sensuDelete %s", err)
					output := fmt.Sprintf("Sorry. Request Failed: %s", err)
					return output
				}
				count++

				// s += fmt.Sprintf("Check Name: %s, Entity: %s, Namespace: %s\n", n.Check.Name, n.Entity.Name, n.Check.Namespace)
			}
			s += fmt.Sprintf("Number of Events deleted: %v \n", count)
			logLocal.Debug("Number of ", command["resource"], " deleted: ", count)

		case "checks":
			result := new([]types.Check)
			_ = json.Unmarshal(data, &result)
			s += fmt.Sprintf("Total of %s found: %v \n", command["resource"], len(*result))
			logLocal.Debug("Total of ", command["resource"], " found: ", len(*result))
			count := 0
			for _, n := range *result {
				if !utils.SearchLabels(k, v, n.Labels) {
					continue
				}
				localSensuURL := sensuURLGenerator(command["resource"], command["namespace"], n.Name)
				err = sensu.SensuDelete(localSensuURL)
				if err != nil {
					logLocal.Errorf("delete verb requestSensu using sensuPost with PUT %s", err)
					output := fmt.Sprintf("Sorry. Request Failed: %s", err)
					return output
				}
				count++

				// s += fmt.Sprintf("Check Name: %s, Entity: %s, Namespace: %s\n", n.Check.Name, n.Entity.Name, n.Check.Namespace)
			}
			s += fmt.Sprintf("Number of Checks deleted: %v \n", count)
			logLocal.Debug("Number of ", command["resource"], " deleted: ", count)

		}

		// end switch resource

	}

	logLocal.Debugf("requestSensu Sensu API Response %s", s)
	return s
}

//sensuURLGenerator func create final sensu url
func sensuURLGenerator(resource string, namespace string, name string) string {
	var verb = []string{"all", "get", "execute", "silence"}
	sensuScheme := config.Values.SensuAPIScheme
	if config.Values.CACertificate != "Absent" || config.Values.SensuAPIScheme != "http" {
		sensuScheme = "https"
	}
	var basicURI string
	switch resource {
	case "health":
		basicURI = string("health")
	case "silenced":
		basicURI = fmt.Sprintf("api/core/v2/namespaces/%s/silenced/%s", namespace, name)
	case "execute":
		basicURI = fmt.Sprintf("api/core/v2/namespaces/%s/checks/%s/execute", namespace, name)
	case "namespaces":
		basicURI = "api/core/v2/namespaces"
	default:
		if !utils.StringInSlice(name, verb) {
			basicURI = fmt.Sprintf("api/core/v2/namespaces/%s/%s/%s", namespace, resource, name)
		} else {
			basicURI = fmt.Sprintf("api/core/v2/namespaces/%s/%s", namespace, resource)
		}
	}

	sensuBase := fmt.Sprintf("%s://%s/%s", sensuScheme, config.Values.SensuAPI, basicURI)
	return sensuBase
}

func checkSensuAPIConfig() bool {
	if config.Values.SensuAPI == "Absent" {
		return false
	}
	// all auth methods are disabled
	if config.Values.SensuAPIToken == "Absent" && config.Values.SensuUser == "Absent" && config.Values.SensuPassword == "Absent" {
		return false
	}
	// sensubot configured to use user and password
	if config.Values.SensuUser != "Absent" && config.Values.SensuPassword != "Absent" && config.Values.SensuAPIToken == "Absent" {
		return true
	}
	// sensubot configuer to use api token
	if config.Values.SensuAPIToken != "Absent" && config.Values.SensuUser == "Absent" && config.Values.SensuPassword == "Absent" {
		return true
	}
	// always return false
	return false
}
