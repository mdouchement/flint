package file

import (
	"path/filepath"
	"strings"

	"github.com/bloom42/flint/lint"
)

type LowerCaseExt struct{}

func (r LowerCaseExt) Apply(file lint.File) []lint.Issue {
	issues := []lint.Issue{}

	if file.IsDir {
		return issues
	}

	ext := filepath.Ext(file.Name)
	if strings.ToLower(ext) != ext {
		issue := lint.Issue{
			Rule:    r.Name(),
			Message: "Extension should be lower cased",
		}
		issues = append(issues, issue)
	}

	return issues
}

func (_ LowerCaseExt) Name() string {
	return "file/lower_case_ext"
}
