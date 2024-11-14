package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:     "power",
	Short:   "Turn your hot tub on or off",
	Long:    `Turn your BestWay hot tub on or off`,
	PreRunE: requireAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no power state provided")
		}

		if len(args) > 1 {
			return fmt.Errorf("too many arguments")
		}

		if args[0] != "on" && args[0] != "off" {
			return fmt.Errorf("invalid power state [on|off]: %s", args[0])
		}

		power := args[0] == "on"

		config := getConfig()
		bw := NewBestWay()
		bw.SetToken(config.Token)
		_, err := bw.SetPower(config.DeviceID, power)
		if err != nil {
			return fmt.Errorf("error setting power: %w", err)
		}

		fmt.Println(
			lipgloss.NewStyle().Foreground(colorWhite).Render("Power set to ") +
				lipgloss.NewStyle().Bold(true).Foreground(colorGreen).Render(fmt.Sprintf("%t", power)),
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
	rootCmd.AddCommand(powerCmd)
}
