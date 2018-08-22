package main

import (
	"fmt"
	"os"

	"github.com/astrocorp42/flint/commands"
)

func fail(err string) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	/*config, err := getConfig()
	//formatter, err := getFormatter(config.Formatter)

	fmt.Println(config)
	fmt.Println(err)
	/*issuesc, errc := linter.Lint(config)
	outputc, errc := formatter.Format(issuesc)


	exitCode := 0
	for out := range outputc {
		fmt.Println(out)
	}


	os.Exit(exitCode)
	*/
}
