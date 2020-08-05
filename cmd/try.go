package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand()
}

var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "Try and possibly fail at something",
	// return an error to the caller of a command
	RunE: func(cmd *cobra.Command, args []string) error {
		err := try()
		if err != nil {
			return err
		}
		return nil
	},
}

func try() error {
	return errors.New("try error")
}
