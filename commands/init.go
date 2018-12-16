package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bloom42/flint/config"
	"github.com/bloom42/sane-go"
	"github.com/spf13/cobra"
)

var initForce bool

func init() {
	RootCmd.AddCommand(InitCmd)
	InitCmd.PersistentFlags().BoolVar(&initForce, "force", false, "Force and override an existing .flint.sane file")
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init flint by creating a .flint.sane configuration file",
	Long:  "Create a flint configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		configFile := config.FindConfigFile()
		var err error

		if configFile != "" && initForce == false {
			fmt.Fprintf(os.Stderr, "A configuration file already exists (%s), use --force to override\n", configFile)
			os.Exit(3)
		}

		conf := config.Default()
		filePath := config.DefaultConfigurationFileName
		data, err := sane.Marshal(conf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(3)
		}

		err = ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(3)
		}
	},
}
