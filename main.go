package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: slack-luxafor <color>\nExample: slack-luxafor red")
		os.Exit(1)
	}

	status := os.Args[1]

	if err := updateSlackStatus(status); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := updateStatusLight(status); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
