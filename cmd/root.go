package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hi",
		Short: "Hi hager!",
	}
)

func init() {
	rootCmd.AddCommand(displayCmd())

}
