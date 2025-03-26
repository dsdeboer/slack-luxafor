/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package utils

import (
	"github.com/slack-go/slack"
)

func UpdateSlackStatus(status string) error {
	value, err := MustGetenv("LUXAFOR_SLACK_API_TOKEN")
	if err != nil {
		return err
	}
	api := slack.New(value)
	return api.SetUserCustomStatus(status, ":speech_balloon:", 0)
}
