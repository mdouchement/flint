package lint

import (
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type Linter struct {
	ExitCode int32
}

func isIgnored(relativePath string, isDir bool, config *Config) bool {
	if isDir {
		// check if is an ignored directory
		for _, regex := range config.IgnoreDirectories {
			if regex.MatchString(relativePath) {
				return true
			}
		}
	} else {
		// check if is an ignored file
		for _, regex := range config.IgnoreFiles {
			if regex.MatchString(relativePath) {
				return true
			}
		}
	}

	return false
}

func (linter *Linter) Lint(config Config) (<-chan Issue, <-chan error) {
	issuesc := make(chan Issue)
	errorsc := make(chan error)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				errorsc <- err
				return nil
			}

			if path == "." || path == ".." {
				return nil
			}

			relativePath, err := filepath.Rel(config.BaseDir, filepath.Join(config.WorkingDir, path))
			if err != nil {
				errorsc <- err
				return nil
			}

			name := info.Name()
			file := File{
				Path:         path,
				Name:         name,
				Ext:          filepath.Ext(name),
				IsDir:        info.IsDir(),
				RelativePath: relativePath,
			}

			if isIgnored(relativePath, file.IsDir, &config) {
				return nil
			}

			// start a new goroutine for each file
			wg.Add(1)
			go func() {
				defer wg.Done()
				issues := lintFile(file, config, errorsc)
				for _, issue := range issues {
					if linter.ExitCode == 0 {
						atomic.StoreInt32(&linter.ExitCode, int32(config.WarningCode))
					}
					if int(linter.ExitCode) != config.ErrorCode {
						if c, ok := config.RulesConfig[issue.RuleName]; ok && c.Severity == SeverityError {
							atomic.StoreInt32(&linter.ExitCode, int32(config.ErrorCode))
						}
					}
					issuesc <- issue
				}
			}()
			return nil
		})
		if err != nil {
			errorsc <- err
		}
	}()

	go func() {
		wg.Wait()
		close(issuesc)
		close(errorsc)
	}()

	return issuesc, errorsc
}

func lintFile(file File, config Config, errorsc <-chan error) []Issue {
	foundIssues := []Issue{}

	for _, currentRule := range config.Rules {
		foundIssues = append(foundIssues, currentRule.Apply(file)...)
	}

	return foundIssues
}

func NewLinter() *Linter {
	return &Linter{}
}
