package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logs out the current user",
	Long:  `Logs out the current user`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("device_id", "")
		viper.Set("token", "")
		viper.Set("token_expires", 0)
		viper.Set("user_id", "")
		viper.Set("username", "")
		viper.Set("password", "")
		viper.WriteConfig()

		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Logged out"))
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
