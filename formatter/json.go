package formatter

import (
	"encoding/json"

	"github.com/astrocorp42/flint/lint"
)

type JSON struct{}

type jsonOutput struct {
	Issues  []lint.Issue `json:"issues"`
	Summary jsonSummary  `json:"summary"`
}

type jsonSummary struct {
	Errors   map[string]uint64 `json:"errors"`
	Warnings map[string]uint64 `json:"warnings"`
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
			Summary: jsonSummary{Errors: map[string]uint64{}, Warnings: map[string]uint64{}},
		}
		for issue := range issuesc {
			output.Issues = append(output.Issues, issue)
			if issue.Severity == lint.SeverityError {
				output.Summary.Errors[issue.Rule] += 1
			} else {
				output.Summary.Warnings[issue.Rule] += 1
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
