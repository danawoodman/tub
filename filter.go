package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Turn the filter on or off",
	Long:  `Turn the filter on or off.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no filter state provided")
		}

		if len(args) > 1 {
			return fmt.Errorf("too many arguments")
		}

		if args[0] != "on" && args[0] != "off" {
			return fmt.Errorf("invalid filter state [on|off]: %s", args[0])
		}

		filter := args[0] == "on"

		config := getConfig()
		bw := NewBestWay()
		bw.SetToken(config.Token)
		_, err := bw.SetFilter(config.DeviceID, filter)
		if err != nil {
			return fmt.Errorf("error setting filter: %w", err)
		}

		fmt.Println(
			lipgloss.NewStyle().Foreground(colorWhite).Render("Filter set to ") +
				lipgloss.NewStyle().Bold(true).Foreground(colorGreen).Render(fmt.Sprintf("%t", filter)),
		)

		time.Sleep(100 * time.Millisecond)

		status, err := bw.GetDeviceStatus(config.DeviceID)
		if err != nil {
			return fmt.Errorf("error getting status: %w", err)
		}

		renderStatus(status)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
}
