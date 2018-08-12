package rule

import (
	"strings"

	"github.com/z0mbie42/flint/lint"
)

type NoDirDot struct{}

func (r NoDirDot) Apply(file lint.File) []lint.Issue {
	issues := []lint.Issue{}

	if file.IsDir != true {
		return issues
	}
	dotCount := strings.Count(file.Name, ".")

	if dotCount >= 1 {
		issue := lint.Issue{
			File:         file,
			RuleName:     r.Name(),
			Explaination: "Unexpected '.' in directory name",
		}
		issues = append(issues, issue)
	}

	return issues
}

func (_ NoDirDot) Name() string {
	return "no_dir_dot"
}
