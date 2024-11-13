package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tub CLI",
	Long:  `Print the version number of tub CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"%s %s\n",
			lipgloss.NewStyle().
				Bold(true).
				Foreground(colorGreen).
				Render(version),
			lipgloss.NewStyle().
				Foreground(colorGray).
				Render(fmt.Sprintf("(commit %s)", lipgloss.NewStyle().Foreground(colorTeal).Render(commit))),
		)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
