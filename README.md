SensuBot
========

SensuBot can receive messages from Slack and/or Telegram. It can answer simple commands like get, execute and silence.

It can list almost all resources available in Sensu: assets, checks, entities, events, namespaces, mutators, filters, handlers, hooks and health.

# Build

```sh
go build
```

# Create token api in sensu backend

```sh
sensuctl user create sensubot --password "LONGPASSWORD"

sensuctl cluster-role-binding create sensubot-rolebinding --cluster-role=cluster-admin --user=sensubot

sensuctl api-key grant sensubot

```

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
  --from-file=sensuCACertificate=./ca.pem \
  -n sensubot --dry-run -o yaml > sensubot-secret.yaml
```

## Deployment

```sh
kubectl create ns sensubot
kubectl apply -f sensubot-secrets.yaml
kubectl apply -f k8s-deployment.yaml
```

If your sensu backend api use https, don't forgot to add CA certificate into secrets and change `SENSUBOT_API_SCHEME` to use https.
