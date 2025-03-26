package main

import (
	"fmt"
	"os"
	"strconv"
)

type LED byte
type COMMAND byte

// Commands recognized by the Luxafor.
const (
	static COMMAND = 1
)

const (
	All LED = 255
)

func MustGetenv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("missing environment variable '%s'", key)
	}
	return val, nil
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
