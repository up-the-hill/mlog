package cmd

import (
	"fmt"
	"up-the-hill/mlog/config"
	"up-the-hill/mlog/utils"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [musing]",
	Short: "Add a new musing",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := config.LoadConfig()
		musingPath := c.MusingsFile
		entry := args[0]
		utils.AppendEntry(musingPath, entry)
		fmt.Printf("Logged \"%s\"\n", entry)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
