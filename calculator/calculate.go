package calculator

import (
	"fmt"
	"math"
	"time"
)

const targetSteps = 5_000_000
const daysInYear = 366

// Calculate calculates the current statistics with the given current steps
func Calculate(currentSteps int) string {

	targetStepsPerDay := targetSteps / daysInYear

	currentDay := dayOfTheYear(time.Now())

	currentTarget := targetStepsPerDay * currentDay

	diff := currentSteps - currentTarget
	relativeDiff := (float64(diff) / float64(currentTarget)) * 100.0

	var plus string
	if diff > 0 {
		plus = "+"
	}

	return fmt.Sprintf("\nCurrent target:\t%d\n"+
		"Current steps:\t%d\n"+
		"Current diff:\t%d\n"+
		"Relative diff:\t%s%.2f%%\n", currentTarget, currentSteps, diff, plus, relativeDiff)
}

func dayOfTheYear(now time.Time) int {
	beginningOfYear := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	timePassed := now.Sub(beginningOfYear)
	daysPassed := int(math.Ceil(timePassed.Hours() / 24))
	return daysPassed
}
