package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hi",
		Short: "Hi hager!",
	}
)

func Init() {
	rootCmd.AddCommand(displayCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
