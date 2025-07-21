package cmd

import (
	"up-the-hill/mlog/config"
	"up-the-hill/mlog/utils"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export all musings to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.LoadConfig()
		utils.ExportMusings(config.MusingsFile, config.ExportPath)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
