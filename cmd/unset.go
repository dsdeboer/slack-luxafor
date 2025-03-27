/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package cmd

import (
	"fmt"

	"github.com/dsdeboer/slack-luxafor-utils"

	"github.com/spf13/cobra"
)

// unsetCmd represents the unset command
var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Sets the status to unset and your light to green",
	Long:  `Quickly set your status to unset and your light to green.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.UpdateSlackStatus("off", ""); err != nil {
			fmt.Println(err)
		}
		if err := utils.UpdateStatusLight("white"); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	statusCmd.AddCommand(unsetCmd)
}
