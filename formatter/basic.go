package formatter

import (
	"fmt"

	"github.com/bloom42/flint/lint"
)

type Basic struct{}

func (Basic) Name() string {
	return "basic"
}

func (formatter Basic) Format(filec <-chan lint.File) (<-chan string, <-chan error) {
	errorsc := make(chan error)
	outputc := make(chan string)

	go func() {
		for file := range filec {
			for _, issue := range file.Issues {
				outputc <- fmt.Sprintf("%s: [%s] %s", file.Path, issue.Rule, issue.Message)
			}
		}
		close(outputc)
		close(errorsc)
	}()

	return outputc, errorsc
}
