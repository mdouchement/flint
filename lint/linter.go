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

func (linter *Linter) Lint(config Config, loadedRules []Rule) (<-chan Issue, <-chan error) {
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
			name := info.Name()
			file := File{
				Path:  path,
				Name:  name,
				Ext:   filepath.Ext(name),
				IsDir: info.IsDir(),
			}

			// start a new goroutine for each file
			wg.Add(1)
			go func() {
				defer wg.Done()
				issues := file.lint(config, loadedRules, errorsc)
				for _, issue := range issues {
					if linter.ExitCode == 0 {
						atomic.StoreInt32(&linter.ExitCode, int32(config.WarningCode))
					}
					if int(linter.ExitCode) != config.ErrorCode {
						if c, ok := config.Rules[issue.RuleName]; ok && c.Severity == SeverityError {
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

func NewLinter() *Linter {
	return &Linter{}
}
