package commands

import (
	"fmt"
	"os"
	"strings"
	"sync"

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

// MergeErrors merges multiple channels of errors.
// Based on https://blog.golang.org/pipelines.
func mergeErrors(cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup

	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan error, len(cs))

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
	output := func(c <-chan error) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines
	// are done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func waitForPipeline(errs ...<-chan error) error {
	errc := mergeErrors(errs...)
	for err := range errc {
		if err != nil {
			return err
		}
	}
	return nil
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

		exitCode, err := runPipeline(conf, choosenFormatter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			os.Exit(3)
		}

		os.Exit(int(exitCode))
	},
}

func runPipeline(config lint.Config, formatter lint.Formatter) (int32, error) {
	var errcList []<-chan error

	// create the pipeline
	linter := lint.NewLinter(config)

	filesc, errc := linter.Lint()
	errcList = append(errcList, errc)

	outputc, errc := formatter.Format(filesc)
	errcList = append(errcList, errc)

	errc = write(outputc)
	errcList = append(errcList, errc)

	err := waitForPipeline(errcList...)
	if err != nil {
		return 3, err
	}

	return linter.ExitCode, nil
}

func write(lines <-chan string) <-chan error {
	ret := make(chan error)

	go func() {
		for outputline := range lines {
			fmt.Println(outputline)
		}
		close(ret)
	}()

	return ret
}
