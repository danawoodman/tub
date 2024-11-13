package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type (
	BestwayConfig struct {
		Token     string
		Username  string
		Password  string
		UserID    string
		DeviceID  string
		ExpiresAt time.Time
		Expired   bool
	}
)

func getConfig() *BestwayConfig {
	deviceID := viper.GetString("device_id")
	token := viper.GetString("token")
	username := viper.GetString("username")
	password := viper.GetString("password")
	userID := viper.GetString("user_id")
	expires := viper.GetInt64("token_expires")
	expiresAt := time.Unix(expires, 0)
	expired := expiresAt.Before(time.Now())

	return &BestwayConfig{
		Token:     token,
		Username:  username,
		Password:  password,
		UserID:    userID,
		DeviceID:  deviceID,
		ExpiresAt: expiresAt,
		Expired:   expired,
	}
}

func requireAuth(cmd *cobra.Command, args []string) error {
	config := getConfig()

	if config.Expired {
		return fmt.Errorf("token expired")
	}

	if config.Token == "" {
		return fmt.Errorf("not logged in")
	}

	return nil
}
