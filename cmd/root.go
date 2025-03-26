/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "slack-luxafor",
	Short: "This application manages your Luxafor light and your Slack status",
	Long: `With this application you can set your Slack status and your Luxafor light at the same time.
You can also turn off the light and unset the status.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
