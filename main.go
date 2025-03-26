package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: slack-luxafor <color>\nExample: slack-luxafor red")
		os.Exit(1)
	}

	status := os.Args[1]

	updateSlackStatus(status)
	updateStatusLight(status)
}

func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("missing environment variable '%s'", key))
	}
	return val
}

func updateSlackStatus(status string) {
	api := slack.New(MustGetenv("LUXAFOR_SLACK_API_TOKEN"))
	err := api.SetUserCustomStatus(status, ":speech_balloon:", 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func updateStatusLight(status string) {
	luxs := Enumerate()
	if len(luxs) == 0 {
		fmt.Println("No attached devices. Exiting.")
		os.Exit(0)
	}

	lux := luxs[1]

	r, g, b := color(status)

	err := lux.Solid(r, g, b)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func color(s string) (uint8, uint8, uint8) {
	rgb := [3]string{"0", "0", "0"}
	switch s {
	case "red":
		rgb[0] = "255"
	case "yellow":
		rgb[0] = "255"
		rgb[1] = "255"
	case "green":
		rgb[1] = "255"
	}
	r, _ := strconv.ParseUint(rgb[0], 10, 8)
	g, _ := strconv.ParseUint(rgb[1], 10, 8)
	b, _ := strconv.ParseUint(rgb[2], 10, 8)
	return uint8(r), uint8(g), uint8(b)
}
