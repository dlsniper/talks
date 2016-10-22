#!/usr/bin/env bash

set -ex

export GOOGLE_APPLICATION_CREDENTIALS=${HOME}/account-auth.json

sudo /opt/google-cloud-sdk/bin/gcloud docker push eu.gcr.io/${PROJECT_NAME}/bot
sudo chown -R ubuntu:ubuntu /home/ubuntu/.kube
kubectl rolling-update --namespace=codemo-slack-bot codemo-slack-bot --image=eu.gcr.io/codemo-slack-bot/bot:${CIRCLE_SHA1} --update-period=10s
