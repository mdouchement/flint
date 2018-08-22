package formatter

import (
	"encoding/json"

	"github.com/astrocorp42/flint/lint"
)

type NDJSON struct{}

func (NDJSON) Name() string {
	return "ndjson"
}

func (formatter NDJSON) Format(issuesc <-chan lint.Issue) (<-chan string, <-chan error) {
	errorsc := make(chan error)
	outputc := make(chan string)

	go func() {
		for issue := range issuesc {
			result, err := json.Marshal(issue)
			if err != nil {
				errorsc <- err
				break
			}
			outputc <- string(result)
		}
		close(outputc)
		close(errorsc)
	}()

	return outputc, errorsc
}
