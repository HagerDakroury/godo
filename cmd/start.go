package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hagerdakroury/godo/display"
)

// in minutes
var duaration int
var title string
var startCmd = &cobra.Command{

	Use:   "start",
	Short: "start a pomodoro",
	Long:  "The 'start' command is used to start a new pomodoro session.",
	RunE: func(cmd *cobra.Command, args []string) error {
		display.ProgressBar(duaration, title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	//default 35 mins
	startCmd.PersistentFlags().IntVarP(&duaration, "duration", "d", 35, "godormo session duration")
	startCmd.PersistentFlags().StringVarP(&title, "title", "t", "task", "godormo session duration")

}
