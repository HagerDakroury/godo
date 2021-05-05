package display

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

func ProgressBar(duration int, task string) {
	bar := progressbar.NewOptions(100,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetPredictTime(true),
	)
	bar.Describe(task)


	for i := 0; i < duration; i++ {
		bar.Add(1)
		time.Sleep(60 * time.Second)
	}

}
