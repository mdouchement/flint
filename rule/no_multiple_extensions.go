package rule

import (
	"strings"

	"github.com/z0mbie42/flint/lint"
)

type NoMultipleExtensions struct{}

func (r NoMultipleExtensions) Apply(file lint.File) []lint.Issue {
	dotCount := strings.Count(file.Name, ".")
	issues := []lint.Issue{}

	if dotCount > 1 {
		issue := lint.Issue{
			File:         file,
			RuleName:     r.Name(),
			Explaination: "should not have multiple extensions (multiple . in name)",
		}
		issues = append(issues, issue)
	}

	return issues
}

func (_ NoMultipleExtensions) Name() string {
	return "no_multiple_extensions"
}
