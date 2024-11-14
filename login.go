package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	username string
	password string

	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to your hot tub",
		Long:  `Login to your BestWay hot tub to generate authentication tokens for other requests.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			bw := NewBestWay()

			session, err := bw.Login(username, password)
			if err != nil {
				panic(err)
			}

			devices, err := bw.ListDevices()
			if err != nil {
				return fmt.Errorf("error listing devices: %w", err)
			}

			if len(devices.Devices) == 0 {
				return fmt.Errorf("no devices found")
			}

			if len(devices.Devices) == 1 {
				viper.Set("device_id", devices.Devices[0].ID)
			}

			// TODO: how to choose from multiple devices?

			renderTable([][]string{
				{"Token", session.Token},
			})

			viper.Set("username", username)
			viper.Set("password", password)
			viper.Set("user_id", session.UID)
			viper.Set("token", session.Token)
			viper.Set("token_expires", session.ExpireAt)
			viper.WriteConfig()

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Your BestWay username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Your BestWay password")
}
