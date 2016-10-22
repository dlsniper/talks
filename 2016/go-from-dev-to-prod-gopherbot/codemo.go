package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

var (
	botName    = os.Getenv("CODEMO_SLACK_BOT_NAME")
	slackToken = os.Getenv("CODEMO_SLACK_BOT_TOKEN")
	botVersion = "HEAD"
	slackAPI   = slack.New(slackToken)
)

func main() {
	if slackToken == "" {
		log.Fatal("slack token must be set in the CODEMO_SLACK_BOT_TOKEN environment variable")
	}
	if botName == "" {
		log.Fatal("bot name missing, set it with CODEMO_SLACK_BOT_NAME")
	}

	rtm := slackAPI.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch message := msg.Data.(type) {
			case *slack.MessageEvent:
				handleMessage(message)
			}
		}
	}
}

func handleMessage(event *slack.MessageEvent) {
	if event.BotID != "" || event.User == "" || event.SubType == "bot_message" {
		return
	}

	eventText := strings.ToLower(event.Text)

	if strings.Contains(eventText, "xkcd:compiling") {
		xkcd(event, "https://xkcd.com/303/")
		return
	}

	if strings.Contains(eventText, strings.ToLower(botName)) {
		if strings.Contains(eventText, "thank") ||
			strings.Contains(eventText, "cheers") ||
			strings.Contains(eventText, "hello") ||
			strings.Contains(eventText, "hi") {
			reactToEvent(event, "wave")
			return
		}

		if strings.Contains(eventText, "version") {
			replyVersion(event)
			return
		}
		return
	}
}

func xkcd(event *slack.MessageEvent, imageLink string) {
	params := slack.PostMessageParameters{AsUser: true, UnfurlLinks: true}
	_, _, err := slackAPI.PostMessage(event.Channel, imageLink, params)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
}

func reactToEvent(event *slack.MessageEvent, reaction string) {
	item := slack.ItemRef{
		Channel:   event.Channel,
		Timestamp: event.Timestamp,
	}
	err := slackAPI.AddReaction(reaction, item)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
}

func replyVersion(event *slack.MessageEvent) {
	params := slack.PostMessageParameters{AsUser: true}
	_, _, err := slackAPI.PostMessage(event.User, fmt.Sprintf("My version is: %s", botVersion), params)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
}
