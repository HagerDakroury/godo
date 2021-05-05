package display

import (
	"fmt"
	"os"
	"time"

	"github.com/ermanimer/progress_bar"
)

func ProgressBar(duration int) {
	output := os.Stdout
	schema := "({bar}) ({percent}) ({current} of {total} completed)"
	filledCharacter := "="
	blankCharacter := "-"
	var length float64 = 60
	var totalValue float64 = float64(duration * 2)

	//new progress bar
	pb := progress_bar.NewProgressBar(output, schema, filledCharacter, blankCharacter, length, totalValue)

	err := pb.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for value := 1; value <= 80; value++ {
		time.Sleep(20 * time.Millisecond)
		err := pb.Update(float64(value))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

}
