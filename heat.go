package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var heatCmd = &cobra.Command{
	Use:   "heat",
	Short: "Turn the heat on or off",
	Long:  `Turn the heat on or off.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no heat state provided")
		}

		if len(args) > 1 {
			return fmt.Errorf("too many arguments")
		}

		if args[0] != "on" && args[0] != "off" {
			return fmt.Errorf("invalid heat state [on|off]: %s", args[0])
		}

		heat := args[0] == "on"

		config := getConfig()
		bw := NewBestWay()
		bw.SetToken(config.Token)
		_, err := bw.SetHeat(config.DeviceID, heat)
		if err != nil {
			return fmt.Errorf("error setting heat: %w", err)
		}

		fmt.Println(
			lipgloss.NewStyle().Foreground(colorWhite).Render("Heat set to ") +
				lipgloss.NewStyle().Bold(true).Foreground(colorGreen).Render(fmt.Sprintf("%t", heat)),
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
	rootCmd.AddCommand(heatCmd)
}
