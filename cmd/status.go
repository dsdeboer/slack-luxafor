/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use: "status",
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
