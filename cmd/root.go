package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the basic first command called with the program runs
var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "Hi Hager, lots of things to do today! LET'S GO!",
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
