package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func updateSlackStatus(status string) {
	api := slack.New(MustGetenv("LUXAFOR_SLACK_API_TOKEN"))
	err := api.SetUserCustomStatus(status, ":speech_balloon:", 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
