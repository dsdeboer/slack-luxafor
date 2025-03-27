/*
Copyright © 2025 Duncan de Boer <duncan.deboer@simplicate.nl>
*/
package cmd

import (
	"fmt"
	"github.com/dsdeboer/slack-luxafor/utils"

	"github.com/getlantern/systray"
	"github.com/spf13/cobra"
)

// taskbarCmd represents the taskbar command
var taskbarCmd = &cobra.Command{
	Use:   "taskbar",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		onExit := func() {
			if err := utils.UpdateSlackStatus("off", ""); err != nil {
				fmt.Println(err)
			}
			if err := utils.ShutdownLight(); err != nil {
				fmt.Println(err)
			}
		}

		systray.Run(onReady, onExit)
	},
}

func onReady() {
	systray.SetTitle("💡")
	systray.SetTooltip("Slack Luxafor Status")

	// Create menu items for each status
	available := systray.AddMenuItem("Available", "Set status to available")
	busy := systray.AddMenuItem("Busy", "Set status to busy")
	focus := systray.AddMenuItem("Focus", "Set status to focus")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "Quit the application")

	// Handle menu item clicks
	go func() {
		for {
			select {
			case <-available.ClickedCh:
				{
					if err := utils.UpdateSlackStatus("available", ""); err != nil {
						fmt.Println(err)
					}
					if err := utils.UpdateStatusLight(utils.ConvertStatusToColor("green")); err != nil {
						fmt.Println(err)
					}
				}
			case <-busy.ClickedCh:
				{
					if err := utils.UpdateSlackStatus("busy", "Busy"); err != nil {
						fmt.Println(err)
					}
					if err := utils.UpdateStatusLight(utils.ConvertStatusToColor("yellow")); err != nil {
						fmt.Println(err)
					}
				}
			case <-focus.ClickedCh:
				{
					if err := utils.UpdateSlackStatus("focus", "I'm focussing"); err != nil {
						fmt.Println(err)
					}
					if err := utils.UpdateStatusLight(utils.ConvertStatusToColor("red")); err != nil {
						fmt.Println(err)
					}
				}
			case <-quit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func init() {
	rootCmd.AddCommand(taskbarCmd)
}
