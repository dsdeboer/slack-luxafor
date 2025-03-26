package main

import (
	"github.com/slack-go/slack"
)

func updateSlackStatus(status string) error {
	value, err := MustGetenv("LUXAFOR_SLACK_API_TOKEN")
	if err != nil {
		return err
	}
	api := slack.New(value)
	return api.SetUserCustomStatus(status, ":speech_balloon:", 0)
}
