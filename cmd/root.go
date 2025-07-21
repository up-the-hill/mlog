package cmd

import (
	"fmt"
	"os"
	"up-the-hill/mlog/config"
	"up-the-hill/mlog/ui"

	"github.com/spf13/cobra"

	tea "github.com/charmbracelet/bubbletea"
)

var rootCmd = &cobra.Command{
	Use:   "mlog",
	Short: "mlog is a simple CLI for logging musings",
	Long: `A simple CLI for logging musings, built with Go and Bubble Tea.

It allows you to quickly add new musings, list existing ones, and export them.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.LoadConfig()
		p := tea.NewProgram(ui.InitialModel(config.MusingsFile, config.CharLimit))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

