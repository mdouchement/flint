package rule

import (
	"strings"

	"github.com/z0mbie42/flint/lint"
)

type NoMultiExtensions struct{}

func (r NoMultiExtensions) Apply(file lint.File) []lint.Issue {
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

func (_ NoMultiExtensions) Name() string {
	return "no_multi_extensions"
}
