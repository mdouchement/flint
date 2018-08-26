package lint

import (
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
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

func (linter *Linter) Lint() (<-chan File, <-chan error) {
	filec := make(chan File)
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

			relativePath, err := filepath.Rel(linter.config.BaseDir, filepath.Join(linter.config.WorkingDir, path))
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
				Issues:       []Issue{},
			}

			if isIgnored(relativePath, file.IsDir, &linter.config) {
				return nil
			}

			// start a new goroutine for each file
			wg.Add(1)
			go func() {
				defer wg.Done()
				issues := lintFile(file, linter.config, errorsc)
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
				file.Issues = issues
				filec <- file
			}()
			return nil
		})
		if err != nil {
			errorsc <- err
		}
	}()

	go func() {
		wg.Wait()
		close(filec)
		close(errorsc)
	}()

	return filec, errorsc
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
