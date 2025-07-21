package cmd

import (
	"fmt"
	"up-the-hill/mlog/config"

	"github.com/spf13/cobra"
)

var configCreateCmd = &cobra.Command{
	Use:   "config-create",
	Short: "Create a default config file",
	Run: func(cmd *cobra.Command, args []string) {
		config.CreateConfig()
		fmt.Printf("Created config file!\n")
	},
}

func init() {
	rootCmd.AddCommand(configCreateCmd)
}
