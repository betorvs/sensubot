SensuBot
========

![Go Test](https://github.com/betorvs/sensubot/workflows/Go%20Test/badge.svg)

SensuBot can receive messages from Slack and/or Telegram. It can answer simple commands like get, execute and silence.

It can list almost all resources available in Sensu: assets, checks, entities, events, namespaces, mutators, filters, handlers, hooks and health.

# Environment Variables

* **SENSUBOT_PORT**: default "9090";
* **SENSUBOT_TIMEOUT**: default "15" seconds;
* **SENSUBOT_DEBUG_SENSU_REQUESTS**:  default "false";
* **SENSUBOT_SLASH_COMMAND**: default "/sensubot";
* **SENSUBOT_SLACK_TOKEN**: Please create one in api.slack and configure it (starts with "xoxb-")
* **SENSUBOT_SLACK_SIGNING_SECRET**: Please get from api.slack these secret;
* **SENSUBOT_SLACK_CHANNEL**: For slack, sensuBot needs to have on channel to listen (looks like "MM34AASDD");
* **SENSUBOT_CA_CERTIFICATE**: If you are using private certificates in Sensu Backend, please share CA public certificate here (like /etc/sensu/ca.pem)
* **SENSUBOT_TELEGRAM_TOKEN**: Please, configure your token here;
* **SENSUBOT_TELEGRAM_URL**: If you want to change Telegram API URL, set this. Default: "https://api.telegram.org/bot";
* **SENSUBOT_API_SCHEME**: If your Sensu Backend are using https, change here. Default: "https";
* **SENSUBOT_API_TOKEN**: Sensu Backend API token;
* **SENSUBOT_API_URL**: Sensu Backend API URL (like "sensu-api.sensu.svc.cluster.local:8080")

# Get Configurations

## Create token api in sensu backend

```sh
sensuctl user create sensubot --password "LONGPASSWORD"

sensuctl cluster-role-binding create sensubot-rolebinding --cluster-role=cluster-admin --user=sensubot

sensuctl api-key grant sensubot

```

## Create a App in Slack

### Add feature Slash Command with these parameters:

* Command: `/sensubot`
* Request URL: `https://URL/sensubot/v1/slack`
* Short Description: `Talk with Monitoring System Sensu Go `
* Usage Hint: `[get|execute|silence] RESOURCE NAME NAMESPACE`

### Add feature "Bots" to this bot.

In Oauth permissions add:
* CONVERSATIONS chat:write:bot
* FILES files:write:user

Install these Application in a Channel.

## Create Bot in Telegram

* Search for "botfather";
* Use "/newbot" command;
* Choose your bot name;
* Get Credentials from botFather.

TIP: If you need to share these SensuBot in a Group, ask to make this public in botFather:

* Use command "/setprivacy";
* Past your bot name;
* Choose "DISABLE".

# Deploy sensubot in Kubernetes

## Create sensuBot secrets

```sh
kubectl create secret generic sensubot --from-literal=sensubotApiToken=LONGHASH \
  --from-literal=slackToken=xxx-9X-zxczxczxczxc --from-literal=slackSigningSecret=asdasdasd-asdasdsad-123 \
  -n sensubot --dry-run -o yaml > sensubot-secret.yaml
```

With CA Certiticate
```sh
kubectl create secret generic sensubot --from-literal=sensubotApiToken=LONGHASH \
  --from-literal=slackToken=xxx-9X-zxczxczxczxc --from-literal=slackSigningSecret=asdasdasd-asdasdsad-123 \
  -n sensubot --dry-run -o yaml > sensubot-secret.yaml
 kubectl create secret generic sensu-ca-pem --from-file=sensuca=./ca.pem -n sensubot --dry-run -o yaml > sensu-ca-secret.yaml
```

## Deployment

Before applying [k8s-deployment.yaml](k8s-deployment.yaml) file using kubectl, change it:
- line 28: choose https or http depends how you configure sensu backend api;
- line 32: from sensu-api.sensu.svc.cluster.local:8080 to your sensu backend api URL;
- line 34: change to your Channel in Slack;
- lines 50 and 51: use only if you use create sensu backend certiticates with your own Certificate Authority (CA). 
- line 91: from sensubot.example.local to your domain.

```sh
kubectl create ns sensubot
kubectl apply -f sensubot-secrets.yaml
kubectl apply -f sensu-ca-secret.yaml
kubectl apply -f k8s-deployment.yaml
```

If your sensu backend api use https, don't forgot to add CA certificate into secrets and change `SENSUBOT_API_SCHEME` to use https.

## Example of commands

in Slack:
```
/sensubot get all checks
/sensubot get health

```

in Telegram:
```
@sensu_example_bot get all checks
```
or directly messages:
```
get all checks
```

# Build

```sh
go build
```

## Test and coverage

Run the tests

```sh 
TESTRUN=true go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```sh
golangci-lint run
```


# references

## Golang Spell
The project was initialized using [Golang Spell](https://github.com/golangspell/golangspell).

## Architectural Model
The Architectural Model adopted to structure the application is based on The Clean Architecture.
Further details can be found here: [The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and in the Clean Architecture Book.