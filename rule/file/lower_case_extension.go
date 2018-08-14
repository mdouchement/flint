package file

import (
	"path/filepath"
	"strings"

	"github.com/z0mbie42/flint/lint"
)

type LowerCaseExtension struct{}

func (r LowerCaseExtension) Apply(file lint.File) []lint.Issue {
	issues := []lint.Issue{}

	if file.IsDir {
		return issues
	}

	ext := filepath.Ext(file.Name)
	if strings.ToLower(ext) != ext {
		issue := lint.Issue{
			File:         file,
			RuleName:     r.Name(),
			Explaination: "Extension should be lower cased",
		}
		issues = append(issues, issue)
	}

	return issues
}

func (_ LowerCaseExtension) Name() string {
	return "file/lower_case_extension"
}
