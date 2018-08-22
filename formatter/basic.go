package formatter

import (
	"fmt"

	"github.com/astrocorp42/flint/lint"
)

type Basic struct{}

func (Basic) Name() string {
	return "basic"
}

func (formatter Basic) Format(issuesc <-chan lint.Issue) (<-chan string, <-chan error) {
	errorsc := make(chan error)
	outputc := make(chan string)

	go func() {
		for issue := range issuesc {
			outputc <- fmt.Sprintf("%s: [%s] %s", issue.File.Path, issue.Rule, issue.Message)
		}
		close(outputc)
		close(errorsc)
	}()

	return outputc, errorsc
}
