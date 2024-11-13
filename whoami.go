package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Prints out the currently logged in user",
	Long:  `Prints out the currently logged in user`,
	Run: func(cmd *cobra.Command, args []string) {
		config := getConfig()

		errStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("196"))

		if config.Username == "" {
			fmt.Println(errStyle.Render("Not logged in"))
			return
		}

		if config.Token == "" {
			fmt.Println(errStyle.Render("No token found"))
			return
		}

		if config.Expired {
			fmt.Println(errStyle.Render("Token expired"))
			return
		}

		if config.UserID == "" {
			fmt.Println(errStyle.Render("No user ID found"))
			return
		}

		if config.Password == "" {
			fmt.Println(errStyle.Render("No password found"))
			return
		}

		renderTable([][]string{
			{"Username", config.Username},
			{"Expires", config.ExpiresAt.Format(time.RFC3339)},
			{"Expired", fmt.Sprintf("%t", config.Expired)},
			{"User ID", config.UserID},
		})
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}
