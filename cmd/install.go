package cmd

import (
	"fmt"
	"github.com/fenger/golibdemos/cmd/subcmd"
	"github.com/spf13/cobra"
)

func init() {

	installCmd.AddCommand(subcmd.GetNginxCmd())

	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install a software",
	Run: func(cmd *cobra.Command, args []string) {
		// runCommand(cmd)
	},
}

func runCommand(cmd *cobra.Command) {
	name := cmd.PersistentFlags().Lookup("name")
	fmt.Println("will install ", name.Value)
}

func GetInstallCmd() *cobra.Command {
	return installCmd
}
