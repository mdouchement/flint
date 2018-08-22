package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/astrocorp42/flint/config"
	"github.com/astrocorp42/flint/lint"
	"github.com/spf13/cobra"
)

var rootFormat string

func init() {
	allFormats := allRootFormats()
	RootCmd.PersistentFlags().StringVarP(&rootFormat, "format", "f", allFormats[0], fmt.Sprintf("Output format. Valid values are [%s]", strings.Join(allFormats, ", ")))
}

func allRootFormats() []string {
	ret := make([]string, len(config.AllFormatters))

	for i, format := range config.AllFormatters {
		ret[i] = format.Name()
	}
	return ret
}

var RootCmd = &cobra.Command{
	Use:   "flint",
	Short: "Flint is a fast and configurable filesystem (file and directory names) linter",
	Long: `A Fast and configurable filesystem (file and directory names) linter.
More information here: https://github.com/astrocorp42/flint`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Get()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(3)
		}

		var choosenFormatter lint.Formatter

		found := false
		for _, format := range config.AllFormatters {
			if rootFormat == format.Name() {
				choosenFormatter = format
				found = true
			}
		}
		if found != true {
			fmt.Fprintf(os.Stderr, "Error: %s is not a valid output format\n", rootFormat)
			os.Exit(3)
		}

		linter := lint.NewLinter()
		issuesc, _ := linter.Lint(conf)
		outputc, errc := choosenFormatter.Format(issuesc)

	Loop:
		for {
			select {
			case outputline, isOpen := <-outputc:
				if isOpen != true {
					break Loop
				}
				fmt.Println(outputline)
			case err, isOpen := <-errc:
				if isOpen && err != nil {
					fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
					os.Exit(3)
				}
			}

		}

		os.Exit(int(linter.ExitCode))
	},
}
