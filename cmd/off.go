/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package cmd

import (
	"github.com/dsdeboer/slack-luxafor/utils"
	"github.com/spf13/cobra"
)

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Shuts down the light",
	Long:  `This command will shut down the light.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.ShutdownLight(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
