package main

import (
	"fmt"
	"os"
)

func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("missing environment variable '%s'", key))
	}
	return val
}
