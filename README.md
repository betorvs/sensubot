SensuBot
========

![Go Test](https://github.com/betorvs/sensubot/workflows/Go%20Test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/betorvs/sensubot/badge.svg?branch=main)](https://coveralls.io/github/betorvs/sensubot?branch=main)

SensuBot was initially design to work with chats (Slack, Telegram) to receive requests from users and send it to Sensu API. It should be stateless and never keep any data. We expand this concept to use multiple integrations and the default integration still be Sensu API. 

Simple Diagram:  

```
┌────────────┐              ┌──────────────┐                ┌────────────────┐
│  Chats     ├─────────────►│ SensuBot     ├───────────────►│ Integrations   │
│            │              │              │                │                │
└────────────┘ ◄────────────┴──────────────┘ ◄──────────────┴────────────────┘
```
Created using [asciiflow][1].

It can list these resources in Sensu: assets, checks, entities, events, namespaces, mutators, filters, handlers, hooks and health.

# Environment Variables

## Basic setup

* **SENSUBOT_PORT**: default "9090";
* **SENSUBOT_TIMEOUT**: default "15" seconds;
* **LOG_LEVEL**:  default "INFO";

## Integrations 

* **SENSUBOT_DEFAULT_INTEGRATION_NAME**: Default integration to connect. Default "sensu".
* **SENSUBOT_API_SCHEME**: If your Sensu Backend are using https, change here. Default: "https";
* **SENSUBOT_API_TOKEN**: Sensu Backend API token;
* **SENSUBOT_API_URL**: Sensu Backend API URL (like "sensu-api.sensu.svc.cluster.local:8080")
* **SENSUBOT_ALERTMANAGER_ENDPOINTS**: Alert Manager integration

## Chat configurations 

* **SENSUBOT_SLASH_COMMAND**: default "/sensubot";
* **SENSUBOT_SLACK_TOKEN**: Please create one in api.slack and configure it (starts with "xoxb-")
* **SENSUBOT_SLACK_SIGNING_SECRET**: Please get from api.slack these secret;
* **SENSUBOT_SLACK_CHANNEL**: For slack, sensuBot needs to have on channel to listen (looks like "MM34AASDD");
* **SENSUBOT_CA_CERTIFICATE**: If you are using private certificates in Sensu Backend, please share CA public certificate here (like /etc/sensu/ca.pem)
* **SENSUBOT_TELEGRAM_TOKEN**: Please, configure your token here;
* **SENSUBOT_TELEGRAM_URL**: If you want to change Telegram API URL, set this. Default: "https://api.telegram.org/bot";
* **SENSUBOT_TELEGRAM_NAME**: Telegram bot name. SensuBot use it to remove sensuBot Name from requests when it is used inside a group chat message.
* **SENSUBOT_GCHAT_PROJECTID**: Google Cloud Project ID (numbers)
* **SENSUBOT_GCHAT_BOT_NAME**: Google Chat Bot name
* **SENSUBOT_GCHAT_SA_PATH**: Path for service account json file

## Chat security options

* **SENSUBOT_BLOCKED_VERBS** : blocked list of verbs (get, execute, silence, delete, resolve)
* **SENSUBOT_BLOCKED_RESOURCES** : blocked list of resources from sensu api
* **SENSUBOT_SLACK_ADMIN_ID_LIST** : User ID from slack, google chat and telegram) allowed to run anything
* **SENSUBOT_TELEGRAM_ADMIN_ID_LIST** : User ID from telegram allowed to run anything
* **SENSUBOT_GCHAT_ADMIN_LIST** : User ID from google chat allowed to run anything


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

### Curl tips

```bash 
curl https://api.telegram.org/bot${my_bot_token}/getWebhookInfo
```

```bash 
curl -X POST https://api.telegram.org/bot${my_bot_token}/setWebhook?url=https://YOUR-URL/sensubot/v1/telegram
```

[Source](https://xabaras.medium.com/setting-your-telegram-bot-webhook-the-easy-way-c7577b2d6f72)

## Create Bot In Google Chat (Hangouts Chat)

* This is a HTTPS Bot
* Create Google Chat Bot [here](https://developers.google.com/chat/how-tos/bots-publish)
  - Avatar URl: `https://docs.sensu.io/images/lizy-logo-a.png`
  - Connection settings: Check `Bot URL` and in `Bot URL` add your external endpoint ending in `https://your-domain.com/sensubot/v1/gchat`
* Create and Download service account JSON [here](https://developers.google.com/chat/how-tos/service-accounts)

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
/sensubot get checks
/sensubot get health

```

in Telegram:
```
@sensu_example_bot get all checks
```
or directly messages:
```
get checks
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


[1]: https://asciiflow.com/#/
[2]: https://petstore.swagger.io/?url=https://raw.githubusercontent.com/prometheus/alertmanager/master/api/v2/openapi.yaml#/