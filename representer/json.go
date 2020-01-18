package representer

import (
	"encoding/json"
	"fmt"

	"github.com/p886/5-million-steps/calculator"
)

// CalculationJSON represents a calculation as JSON
type CalculationJSON struct {
	CurrentTarget int    `json:"current_target"`
	CurrentSteps  int    `json:"current_steps"`
	Diff          int    `json:"diff"`
	RelativeDiff  string `json:"relative_diff"`
}

// AsJSON generates a JSON representation of a Calculation
func AsJSON(result calculator.CalculationResult) ([]byte, error) {
	calcJSON := CalculationJSON{
		CurrentTarget: result.CurrentTarget,
		CurrentSteps:  result.CurrentSteps,
		Diff:          result.Diff,
		RelativeDiff:  fmt.Sprintf("%s%.2f%%", result.Plus, result.RelativeDiff),
	}
	json, err := json.Marshal(&calcJSON)
	return json, err
}
