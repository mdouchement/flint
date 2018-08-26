package formatter

import (
	"encoding/json"

	"github.com/astrocorp42/flint/lint"
)

type NDJSON struct{}

func (NDJSON) Name() string {
	return "ndjson"
}

func (formatter NDJSON) Format(filec <-chan lint.File) (<-chan string, <-chan error) {
	errorsc := make(chan error)
	outputc := make(chan string)

	go func() {
		for file := range filec {
			if len(file.Issues) != 0 {
				result, err := json.Marshal(file)
				if err != nil {
					errorsc <- err
					break
				}
				outputc <- string(result)
			}
		}
		close(outputc)
		close(errorsc)
	}()

	return outputc, errorsc
}
