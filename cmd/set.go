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

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use: "set",
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		// Run the custom validation logic
		if utils.IsValidStatus(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid status specified: %s", args[0])
	},
	Short: "Sets the status",
	Long:  `Given you've set up the slack token and the luxafor device, you can set the status with this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		status := args[0]
		message, _ := cmd.Flags().GetString("message")
		if err := utils.UpdateSlackStatus(status, message); err != nil {
			fmt.Println(err)
		}
		if err := utils.UpdateStatusLight(utils.ConvertStatusToColor(status)); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	statusCmd.AddCommand(setCmd)
	setCmd.Flags().String("message", "", "A message to accompany the status")
}
