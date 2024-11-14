package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Get the status of your hot tub",
	Long:    `Get the status of your BestWay hot tub.`,
	PreRunE: requireAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := viper.GetString("token")

		// todo: if no device ID, list devices and if only one device choose that, otherwise prompt

		deviceID := viper.GetString("device_id")
		if deviceID == "" {
			return fmt.Errorf("no device ID provided, please provide a device ID")
		}

		bw := NewBestWay()
		bw.SetToken(token)

		status, err := bw.GetDeviceStatus(deviceID)
		if err != nil {
			return fmt.Errorf("error getting device status: %s", err)
		}

		viper.Set("device_id", deviceID)
		viper.WriteConfig()

		renderStatus(status)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.PersistentFlags().StringP("device", "d", viper.GetString("device_id"), "The device ID to get the status of")
	viper.BindPFlag("device_id", statusCmd.PersistentFlags().Lookup("device"))
}

func renderStatus(status *BestWayDeviceStatusResponse) {
	power := "Off"
	if status.DeviceState.Power == POWER_ON {
		power = "On"
	}

	unit := getTemperatureUnit(status.DeviceState.TemperatureUnit)

	screenLock := "Off"
	if status.DeviceState.ScreenLock == SCREEN_LOCK {
		screenLock = "On"
	}

	jets := "Off"
	if status.DeviceState.Jets == JETS_HIGH {
		jets = "High"
	} else if status.DeviceState.Jets == JETS_LOW {
		jets = "Low"
	}

	heat := "Off"
	if status.DeviceState.Heat > 0 {
		heat = "On"
	}

	filter := "Off"
	if status.DeviceState.Filter == FILTER_ON {
		filter = "On"
	}

	renderTable([][]string{
		{"Power", power},
		{"Set Temp", fmt.Sprintf("%d %s", status.DeviceState.SetTemperature, unit)},
		{"Current Temp", fmt.Sprintf("%d %s", status.DeviceState.CurrentTemperature, unit)},
		{"Filter", filter},
		{"Heat", heat},
		{"Jets", jets},
		{"Screen Lock", screenLock},
	})
}
