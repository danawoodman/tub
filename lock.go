package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock the tub",
	Long:  `Lock the tub to prevent anyone from using it.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no lock state provided")
		}

		if len(args) > 1 {
			return fmt.Errorf("too many arguments")
		}

		if args[0] != "on" && args[0] != "off" {
			return fmt.Errorf("invalid lock state [on|off]: %s", args[0])
		}

		lock := args[0] == "on"

		config := getConfig()
		bw := NewBestWay()
		bw.SetToken(config.Token)
		_, err := bw.SetScreenLock(config.DeviceID, lock)
		if err != nil {
			return fmt.Errorf("error setting lock: %w", err)
		}

		fmt.Println(
			lipgloss.NewStyle().Foreground(colorWhite).Render("Lock set to ") +
				lipgloss.NewStyle().Bold(true).Foreground(colorGreen).Render(fmt.Sprintf("%t", lock)),
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
	rootCmd.AddCommand(lockCmd)
}
