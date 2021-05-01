package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func displayCmd() *cobra.Command {
	return &cobra.Command{
		Use: "displayT",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println(prettyTime)
			return nil
		},
	}
}
