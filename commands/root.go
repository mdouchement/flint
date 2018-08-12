package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/z0mbie42/flint/config"
)

var RootCmd = &cobra.Command{
	Use:   "flint",
	Short: "Flint is a fast and configurable filesystem (file and directory names) linter",
	Long: `A Fast and configurable filesystem (file and directory names) linter.
More information here: https://github.com/z0mbie42/flint`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.Get()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(config)
	},
}
