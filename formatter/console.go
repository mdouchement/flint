package formatter

import (
	"fmt"
	"strconv"

	"github.com/bloom42/flint/lint"
)

const (
	warningIcon = "⚠"
	errorIcon   = "✘"
	goodIcon    = "✔"
)

const (
	cReset    = 0
	cBold     = 1
	cRed      = 31
	cGreen    = 32
	cYellow   = 33
	cBlue     = 34
	cMagenta  = 35
	cCyan     = 36
	cGray     = 37
	cDarkGray = 90
)

func colorize(str string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, str)
}

type Console struct{}

type consoleOutput struct {
	Issues  []lint.Issue `json:"issues"`
	Summary summary      `json:"summary"`
}

func (Console) Name() string {
	return "console"
}

func (formatter Console) Format(filec <-chan lint.File) (<-chan string, <-chan error) {
	errorsc := make(chan error)
	outputc := make(chan string)

	go func() {
		summary := summary{Errors: severitySummary{Rules: map[string]uint64{}}, Warnings: severitySummary{Rules: map[string]uint64{}}}
		finalIcon := goodIcon
		var totalIssues uint64
		errorsStr := "error"
		warningsStr := "warning"
		issuesStr := "issue"
		finalColor := cGreen

		for file := range filec {
			if file.IsDir {
				summary.TotalDirectories += 1
			} else {
				summary.TotalFiles += 1
			}

			for _, issue := range file.Issues {
				totalIssues += 1
				var icon string
				color := cYellow
				if issue.Severity == lint.SeverityError {
					summary.Errors.Total += 1
					summary.Errors.Rules[issue.Rule] += 1
					color = cRed
					icon = errorIcon
				} else {
					summary.Warnings.Total += 1
					summary.Warnings.Rules[issue.Rule] += 1
					icon = warningIcon
				}
				outputc <- fmt.Sprintf("%s [%s] %s\n  %s\n", colorize(icon, color), colorize(issue.Rule, cCyan), issue.Message, file.RelativePath)
			}
		}

		if summary.Errors.Total != 0 {
			finalIcon = errorIcon
			finalColor = cRed
		} else if summary.Warnings.Total != 0 {
			finalIcon = warningIcon
			finalColor = cYellow
		}
		if totalIssues != 1 {
			issuesStr += "s"
		}
		if summary.Errors.Total != 1 {
			errorsStr += "s"
		}
		if summary.Warnings.Total != 1 {
			warningsStr += "s"
		}

		outputc <- colorize(fmt.Sprintf(
			"%s %d %s (%d %s, %d %s)", finalIcon,
			totalIssues, issuesStr,
			summary.Warnings.Total, errorsStr,
			summary.Errors.Total, warningsStr,
		), finalColor)

		if totalIssues != 0 {
			if summary.Warnings.Total != 0 {
				outputc <- colorize("Warnings:", cYellow)
				for rule, count := range summary.Warnings.Rules {
					outputc <- fmt.Sprintf("  %s [%s]", colorize(strconv.FormatInt(int64(count), 10), cYellow), colorize(rule, cCyan))
				}
			}
			if summary.Errors.Total != 0 {
				outputc <- colorize("Errors:", cRed)
				for rule, count := range summary.Errors.Rules {
					outputc <- fmt.Sprintf("  %s [%s]", colorize(strconv.FormatInt(int64(count), 10), cRed), colorize(rule, cCyan))
				}
			}
		}

		close(outputc)
		close(errorsc)
	}()

	return outputc, errorsc
}
