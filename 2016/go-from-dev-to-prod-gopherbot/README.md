# Go Slack bot in 1 hour or less

```

// paste output in GCLOUD_SERVICE_KEY env var in CircleCI
base64 <gcloud-service-key.json>

gcloud auth activate-service-account --key-file ./gcloud-service-key.json
gcloud config set project codemo-slack-bot


gcloud container clusters create codemo-slack-bot \
    --zone europe-west1-c \
    --additional-zones=europe-west1-d,europe-west1-b \
    --num-nodes=1 \
    --local-ssd-count=0 \
    --machine-type=f1-micro \
    --disk-size=10

kubectl create namespace codemo-slack-bot

cp secrets.yaml.template secrets.yaml
echo `echo -n 'slackTokenHere' | base64` >> secrets.yaml
kubectl create -f ./secrets.yaml --namespace=codemo-slack-bot
kubectl create -f ./rc.yaml --namespace=codemo-slack-bot
```
