package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "able",
		Short: "A generator for Cobra based Applications",
		Long:  `Able is a CLI library for Go that empowers applications`,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("able command run")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.SetVersionTemplate("Able version {{.Version}}\n")

}

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}
