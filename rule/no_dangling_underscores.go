package rule

import (
	"strings"

	"github.com/z0mbie42/flint/lint"
)

type NoDanglingUnderscores struct{}

func (r NoDanglingUnderscores) Apply(file lint.File) []lint.Issue {
	parts := strings.Split(strings.TrimSuffix(file.Name, file.Ext), ".")
	issues := []lint.Issue{}

	for _, part := range parts {
		if strings.Trim(part, "_") != part {
			issue := lint.Issue{
				File:         file,
				RuleName:     r.Name(),
				Explaination: "Unexpected dangling '_'",
			}
			issues = append(issues, issue)
		}
	}

	return issues
}

func (_ NoDanglingUnderscores) Name() string {
	return "no_dangling_underscores"
}
