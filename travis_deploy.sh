#!/bin/bash
# basic script to deploy docker in dockerhub with commit and release tags
set -e


IMAGE="${1}"
TAG="${2}"
LATEST="${3}"

docker build -f Dockerfile -t ${IMAGE}:${TAG} .

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push ${IMAGE}:$TAG
if [ "${LATEST}" = "true" ]; then 
  docker tag ${IMAGE}:$TAG ${IMAGE}:latest
  docker push ${IMAGE}:latest
fi