package cmd

import (
	"fmt"
	"up-the-hill/mlog/config"
	"up-the-hill/mlog/utils"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all musings",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.LoadConfig()
		musingPath := config.MusingsFile

		for _, m := range utils.GetEntries(musingPath) {
			dateStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
			fmt.Printf("%s %s\n", dateStyle.Render(m.Date.Format("2006-01-02")), m.Musing)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
