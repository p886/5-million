package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const targetSteps = 5_000_000
const daysInYear = 366

func main() {
	currentSteps, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatalf("Error converting input steps: %v", err)
	}
	targetStepsPerDay := targetSteps / daysInYear

	currentDay, err := dayOfTheYear(time.Now())
	if err != nil {
		log.Fatalf("Error converting input steps: %v", err)
	}

	currentTarget := targetStepsPerDay * currentDay

	diff := currentSteps - currentTarget
	relativeDiff := (float64(diff) / float64(currentTarget)) * 100.0

	fmt.Printf("\nCurrent target:\t%d\n"+
		"Current steps:\t%d\n"+
		"Current diff:\t%d\n"+
		"Relative diff:\t%.2f%%\n", currentTarget, currentSteps, diff, relativeDiff)

}

func dayOfTheYear(now time.Time) (int, error) {
	berlinTimeZone, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return 0, err
	}
	beginningOfYear := time.Date(2020, 1, 1, 0, 0, 0, 0, berlinTimeZone)
	timePassed := now.Sub(beginningOfYear)
	daysPassed := int(math.Ceil(timePassed.Hours() / 24))
	return daysPassed, nil
}
