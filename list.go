package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your hottubs",
	Long:    `List your BestWay hottubs (and other devices)`,
	PreRunE: requireAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := getConfig()

		bw := NewBestWay()
		bw.SetToken(config.Token)

		devices, err := bw.ListDevices()
		if err != nil {
			return fmt.Errorf("error listing devices: %w", err)
		}

		for _, device := range devices.Devices {
			renderTable(
				[][]string{
					{"Name", device.ProductName},
					{"ID", device.ID},
					{"Online", fmt.Sprintf("%t", device.Online)},
					{"Product Name", device.ProductName},
				},
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
