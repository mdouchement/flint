package formatter

import (
	"encoding/json"

	"github.com/astrocorp42/flint/lint"
)

type JSON struct{}

type jsonOutput struct {
	FilesWithIssues []lint.File `json:"files_with_issues"`
	Summary         summary     `json:"summary"`
}

func (JSON) Name() string {
	return "json"
}

func (formatter JSON) Format(filec <-chan lint.File) (<-chan string, <-chan error) {
	errorsc := make(chan error)
	outputc := make(chan string)

	go func() {
		output := jsonOutput{
			FilesWithIssues: []lint.File{},
			Summary:         summary{Errors: severitySummary{Rules: map[string]uint64{}}, Warnings: severitySummary{Rules: map[string]uint64{}}},
		}
		for file := range filec {
			if len(file.Issues) != 0 {
				output.FilesWithIssues = append(output.FilesWithIssues, file)
				for _, issue := range file.Issues {
					if issue.Severity == lint.SeverityError {
						output.Summary.Errors.Total += 1
						output.Summary.Errors.Rules[issue.Rule] += 1
					} else {
						output.Summary.Warnings.Total += 1
						output.Summary.Warnings.Rules[issue.Rule] += 1
					}
				}
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
