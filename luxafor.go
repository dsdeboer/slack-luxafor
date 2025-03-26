package main

import (
	"fmt"
	"os"
)

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
