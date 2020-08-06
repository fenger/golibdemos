package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Able",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Able Anything When Development v0.1 -- HEAD")
	},
}
