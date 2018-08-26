package dir

import (
	"strings"

	"github.com/astrocorp42/flint/lint"
)

type NoDot struct{}

func (r NoDot) Apply(file lint.File) []lint.Issue {
	issues := []lint.Issue{}

	if file.IsDir != true {
		return issues
	}
	dotCount := strings.Count(file.Name, ".")

	if dotCount >= 1 {
		issue := lint.Issue{
			Rule:    r.Name(),
			Message: "Unexpected '.' in directory name",
		}
		issues = append(issues, issue)
	}

	return issues
}

func (_ NoDot) Name() string {
	return "dir/no_dot"
}
