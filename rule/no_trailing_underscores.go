package rule

import (
	"strings"

	"github.com/bloom42/flint/lint"
)

type NoTrailingUnderscores struct{}

func (r NoTrailingUnderscores) Apply(file lint.File) []lint.Issue {
	parts := strings.Split(strings.TrimSuffix(file.Name, file.Ext), ".")
	issues := []lint.Issue{}

	for _, part := range parts {
		if strings.TrimRight(part, "_") != part {
			issue := lint.Issue{
				Rule:    r.Name(),
				Message: "Unexpected trailing '_'",
			}
			issues = append(issues, issue)
		}
	}

	return issues
}

func (_ NoTrailingUnderscores) Name() string {
	return "no_trailing_underscores"
}
