package main

import (
	"fmt"
)

func updateStatusLight(status string) error {
	luxs := Enumerate()
	if len(luxs) == 0 {
		fmt.Println("No attached devices. Exiting.")
		return nil
	}

	lux := luxs[1]

	r, g, b := color(status)

	return lux.Solid(r, g, b)
}
