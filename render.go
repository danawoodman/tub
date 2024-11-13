package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func renderTable(data [][]string) {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(colorPurple)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case col == 0:
				return lipgloss.NewStyle().Foreground(colorGray).Width(16).Padding(0, 1)
			case col == 1:
				return lipgloss.NewStyle().Foreground(colorTeal).Width(34).Padding(0, 1)
			default:
				return lipgloss.NewStyle().Foreground(colorLightGray)
			}
		}).
		Rows(data...)

	fmt.Println(t)
}
