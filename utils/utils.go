/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package utils

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

// MustGetenv retrieves the value of the environment variable named by the key.
func MustGetenv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("missing environment variable '%s'", key)
	}
	return val, nil
}

// IsValidColor checks if the color is valid.
func IsValidColor(s string) bool {
	switch s {
	case "red", "yellow", "green", "white":
		return true
	default:
		return false
	}
}

// ConvertStatusToColor converts a status to a color.
func ConvertStatusToColor(s string) string {
	switch s {
	case "focus":
		return "red"
	case "busy":
		return "yellow"
	case "available":
		return "green"
	default:
		return "white"
	}
}

// IsValidStatus checks if the status is valid.
func IsValidStatus(s string) bool {
	switch s {
	case "focus", "busy", "available", "off":
		return true
	default:
		return false
	}
}
func convertColor(s string) (uint8, uint8, uint8) {
	rgb := [3]string{"0", "0", "0"}
	switch s {
	case "red":
		rgb[0] = "255"
	case "yellow":
		rgb[0] = "255"
		rgb[1] = "255"
	case "green":
		rgb[1] = "255"
	case "white":
		rgb[0] = "255"
		rgb[1] = "255"
		rgb[2] = "255"
	}
	r, _ := strconv.ParseUint(rgb[0], 10, 8)
	g, _ := strconv.ParseUint(rgb[1], 10, 8)
	b, _ := strconv.ParseUint(rgb[2], 10, 8)
	return uint8(r), uint8(g), uint8(b)
}
