package formatter

import (
	"encoding/json"

	"github.com/astrocorp42/flint/lint"
)

type JSON struct{}

type jsonOutput struct {
	Issues  []lint.Issue `json:"issues"`
	Summary summary      `json:"summary"`
}

func (JSON) Name() string {
	return "json"
}

func (formatter JSON) Format(issuesc <-chan lint.Issue) (<-chan string, <-chan error) {
	errorsc := make(chan error)
	outputc := make(chan string)

	go func() {
		output := jsonOutput{
			Issues:  []lint.Issue{},
			Summary: summary{Errors: severitySummary{Rules: map[string]uint64{}}, Warnings: severitySummary{Rules: map[string]uint64{}}},
		}
		for issue := range issuesc {
			output.Issues = append(output.Issues, issue)
			if issue.Severity == lint.SeverityError {
				output.Summary.Errors.Total += 1
				output.Summary.Errors.Rules[issue.Rule] += 1
			} else {
				output.Summary.Warnings.Total += 1
				output.Summary.Warnings.Rules[issue.Rule] += 1
			}
		}
		result, err := json.Marshal(output)
		if err != nil {
			errorsc <- err
		}
		outputc <- string(result)
		close(outputc)
		close(errorsc)
	}()

	return outputc, errorsc
}
