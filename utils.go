package main

import (
	"fmt"
	"os"
)

func MustGetenv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("missing environment variable '%s'", key)
	}
	return val, nil
}
