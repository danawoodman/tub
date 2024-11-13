package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var tempCmd = &cobra.Command{
	Use:     "temp",
	Short:   "Get or set the temperature of your hottub",
	Long:    `Get or set the temperature of your BestWay hottub.`,
	PreRunE: requireAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := getConfig()

		bw := NewBestWay()
		bw.SetToken(config.Token)

		if len(args) <= 0 {
			status, err := bw.GetDeviceStatus(config.DeviceID)
			if err != nil {
				return fmt.Errorf("error getting device status: %s", err)
			}

			temp := getTemperatureUnit(status.DeviceState.TemperatureUnit)
			renderTable(
				[][]string{
					{"Set Temp", fmt.Sprintf("%d %s", status.DeviceState.CurrentTemperature, temp)},
					{"Current Temp", fmt.Sprintf("%d %s", status.DeviceState.CurrentTemperature, temp)},
				},
			)

			return nil
		}

		temp, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid temperature value: %s", args[0])
		}

		resp, err := bw.SetTemp(config.DeviceID, temp)
		if err != nil {
			return fmt.Errorf("error setting temperature: %s", err)
		}

		if resp.StatusCode != http.StatusOK {
			// todo: handle errors like setting temp too high or low
			return fmt.Errorf("error setting temperature: %s", resp.Status)
		}

		time.Sleep(100 * time.Millisecond)

		status, err := bw.GetDeviceStatus(config.DeviceID)
		if err != nil {
			return fmt.Errorf("error getting device status: %s", err)
		}

		fmt.Println(
			lipgloss.NewStyle().Foreground(colorWhite).Render("Temperature set to: ") +
				lipgloss.NewStyle().Foreground(colorGreen).Render(
					fmt.Sprintf("%d %s", temp, getTemperatureUnit(status.DeviceState.TemperatureUnit)),
				),
		)

		renderStatus(status)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(tempCmd)
}

func getTemperatureUnit(unit int) string {
	if unit == 1 {
		return "C°"
	}
	return "F°"
}
