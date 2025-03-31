/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package utils

import (
	"github.com/slack-go/slack"
)

func api() (*slack.Client, error) {
	value, err := MustGetenv("LUXAFOR_SLACK_API_TOKEN")
	if err != nil {
		return nil, err
	}
	return slack.New(value), nil
}
func UpdateSlackStatus(status, message string) error {
	api, err := api()
	if err != nil {
		return err
	}
	switch status {
	case "available", "off":
		if err := api.SetUserPresence("auto"); err != nil {
			return err
		}

		return api.SetUserCustomStatus("", "", 0)
	default:
		emoji := convertStatusToEmoji(status)
		if message == "" {
			message = status
		}
		if message == "none" {
			message = ""
			emoji = ""
		}
		return api.SetUserCustomStatus(message, emoji, 0)
	}
}

func convertStatusToEmoji(status string) string {
	switch status {
	case "focus":
		return ":headphones:"

	default:
		return ":speech_balloon:"
	}
}
