package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func updateSlackStatus(status string) {
	value, err := MustGetenv("LUXAFOR_SLACK_API_TOKEN")
	if err != nil {
		panic(err)
	}
	api := slack.New(value)
	errStatus := api.SetUserCustomStatus(status, ":speech_balloon:", 0)
	if errStatus != nil {
		fmt.Println(errStatus.Error())
		os.Exit(1)
	}

}
