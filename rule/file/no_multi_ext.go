package file

import (
	"strings"

	"github.com/bloom42/flint/lint"
)

type NoMultiExt struct{}

func (r NoMultiExt) Apply(file lint.File) []lint.Issue {
	var name string

	if strings.Index(file.Name, ".") == 0 {
		name = file.Name[1:]
	} else {
		name = file.Name
	}
	dotCount := strings.Count(name, ".")
	issues := []lint.Issue{}

	if file.IsDir {
		return issues
	}

	if dotCount > 1 {
		issue := lint.Issue{
			Rule:    r.Name(),
			Message: "should not have multiple extensions (multiple . in name)",
		}
		issues = append(issues, issue)
	}

	return issues
}

func (_ NoMultiExt) Name() string {
	return "file/no_multi_ext"
}
