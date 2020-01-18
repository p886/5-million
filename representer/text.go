package representer

import "fmt"

import "github.com/p886/5-million-steps/calculator"

// AsText generates a JSON representation of a Calculation
func AsText(calculationResult calculator.CalculationResult) string {
	return fmt.Sprintf("\nCurrent target:\t%d\n"+
		"Current steps:\t%d\n"+
		"Current diff:\t%d\n"+
		"Relative diff:\t%s%.2f%%\n", calculationResult.CurrentTarget,
		calculationResult.CurrentSteps,
		calculationResult.Diff,
		calculationResult.Plus,
		calculationResult.RelativeDiff)
}
