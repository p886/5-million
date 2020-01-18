package calculator

import (
	"math"
	"time"
)

const targetSteps = 5_000_000
const daysInYear = 366

// CalculationResult is the result returned from a calculcation
type CalculationResult struct {
	CurrentTarget int
	CurrentSteps  int
	Diff          int
	Plus          string
	RelativeDiff  float64
}

// Calculate calculates the current statistics with the given current steps
func Calculate(currentSteps int) CalculationResult {

	targetStepsPerDay := targetSteps / daysInYear

	currentDay := dayOfTheYear(time.Now())

	currentTarget := targetStepsPerDay * currentDay

	diff := currentSteps - currentTarget
	relativeDiff := (float64(diff) / float64(currentTarget)) * 100.0

	var plus string
	if diff > 0 {
		plus = "+"
	}

	return CalculationResult{
		CurrentTarget: currentTarget,
		CurrentSteps:  currentSteps,
		Diff:          diff,
		Plus:          plus,
		RelativeDiff:  relativeDiff,
	}
}

func dayOfTheYear(now time.Time) int {
	beginningOfYear := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	timePassed := now.Sub(beginningOfYear)
	daysPassed := int(math.Ceil(timePassed.Hours() / 24))
	return daysPassed
}
