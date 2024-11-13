package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var jetsCmd = &cobra.Command{
	Use:   "jets",
	Short: "Turn the jets on or off",
	Long:  `Turn the jets on or off.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no jets state provided")
		}

		if len(args) > 1 {
			return fmt.Errorf("too many arguments")
		}

		var jets int
		switch args[0] {
		case "off":
			jets = JETS_OFF
		case "low":
			jets = JETS_LOW
		case "high":
			jets = JETS_HIGH
		default:
			return fmt.Errorf("invalid jets state [off|low|high]: %s", args[0])
		}

		config := getConfig()
		bw := NewBestWay()
		bw.SetToken(config.Token)
		_, err := bw.SetJets(config.DeviceID, jets)
		if err != nil {
			return fmt.Errorf("error setting jets: %w", err)
		}

		fmt.Println(
			lipgloss.NewStyle().Foreground(colorWhite).Render("Jets set to ") +
				lipgloss.NewStyle().Bold(true).Foreground(colorGreen).Render(args[0]),
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
	rootCmd.AddCommand(jetsCmd)
}
