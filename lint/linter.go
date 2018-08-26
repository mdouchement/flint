package lint

import (
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/z0mbie42/fswalk"
)

type Linter struct {
	ExitCode int32
	config   Config
}

func isIgnored(relativePath string, isDir bool, config *Config) bool {
	if isDir {
		return config.IgnoreDirectories.MatchString(relativePath)
	}

	return config.IgnoreFiles.MatchString(relativePath)
}

func (linter *Linter) Lint(filesc <-chan fswalk.File) (<-chan File, <-chan error) {
	outc := make(chan File)
	errorsc := make(chan error)

	go func() {
		var wg sync.WaitGroup

		for file := range filesc {
			if file.Path == "." || file.Path == ".." {
				continue
			}

			relativePath, err := filepath.Rel(linter.config.BaseDir, filepath.Join(linter.config.WorkingDir, file.Path))
			if err != nil {
				errorsc <- err
				continue
			}

			fileToLint := File{
				File:         file,
				RelativePath: relativePath,
				Ext:          filepath.Ext(file.Name),
				Issues:       []Issue{},
			}

			if isIgnored(relativePath, file.IsDir, &linter.config) {
				continue
			}

			// start a new goroutine for each file
			wg.Add(1)
			go func() {
				defer wg.Done()
				issues := lintFile(fileToLint, linter.config, errorsc)
				for _, issue := range issues {
					if linter.ExitCode == 0 {
						atomic.StoreInt32(&linter.ExitCode, int32(linter.config.WarningExitCode))
					}
					if c, ok := linter.config.RulesConfig[issue.Rule]; ok && c.Severity == SeverityError {
						issue.Severity = SeverityError
						if int(linter.ExitCode) != linter.config.ErrorExitCode {
							atomic.StoreInt32(&linter.ExitCode, int32(linter.config.ErrorExitCode))
						}
					} else {
						issue.Severity = SeverityWarning
					}
				}
				fileToLint.Issues = issues
				outc <- fileToLint
			}()
		}
		wg.Wait()
		close(outc)
		close(errorsc)
	}()

	return outc, errorsc
}

func lintFile(file File, config Config, errorsc <-chan error) []Issue {
	foundIssues := []Issue{}

	for _, currentRule := range config.Rules {
		foundIssues = append(foundIssues, currentRule.Apply(file)...)
	}

	return foundIssues
}

func NewLinter(config Config) Linter {
	return Linter{config: config}
}
