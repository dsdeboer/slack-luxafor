/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package utils

import (
	"fmt"
)

func UpdateStatusLight(color string) error {
	luxs := Enumerate()
	if len(luxs) == 0 {
		fmt.Println("No attached devices. Exiting.")
		return nil
	}

	lux := luxs[1]

	r, g, b := convertColor(color)

	return lux.Solid(r, g, b)
}

func ShutdownLight() error {
	luxs := Enumerate()
	if len(luxs) == 0 {
		fmt.Println("No attached devices. Exiting.")
		return nil
	}

	lux := luxs[1]

	return lux.Off()
}
