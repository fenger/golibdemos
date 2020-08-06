package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {

	runCmd.PersistentFlags().StringP("daemon", "d", "", "Run something in background")

	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run something. eg: jar or shell.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run command.")
	},
}
