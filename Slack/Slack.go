package Slack

import (
	"github.com/bluele/slack"
)


func SlackMessage(message string, token string, channelName string ) {

	api := slack.New(token)

	err := api.ChatPostMessage(channelName, message, nil)
	if err != nil {
		panic(err)
	}
}


