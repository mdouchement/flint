package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/z0mbie42/flint/config"
)

var initFormat string
var initForce bool

func init() {
	RootCmd.AddCommand(InitCmd)
	InitCmd.PersistentFlags().StringVarP(&initFormat, "format", "f", "toml", "Format of the configuration file. Valid values are [toml, json]")
	InitCmd.PersistentFlags().BoolVar(&initForce, "force", false, "Force and override an existing .flint.(toml|json) file")
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init flint by creating a .flint.(toml|json) configuration file",
	Long:  "Create a flint configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		configFile := config.FindConfigFile()
		var err error

		if configFile != "" && initForce == false {
			fmt.Fprintf(os.Stderr, "A configuration file already exists (%s), use --force to override\n", configFile)
			os.Exit(1)
		}

		if initFormat != "toml" && initFormat != "json" {
			fmt.Fprintf(os.Stderr, "%s is not a valid configuration file format", initFormat)
			os.Exit(1)
		}

		conf := config.Default()
		filePath := config.DefaultConfigurationFileName
		buf := new(bytes.Buffer)

		switch initFormat {
		case "toml":
			err = toml.NewEncoder(buf).Encode(conf)
			filePath += ".toml"
		case "json":
			err = json.NewEncoder(buf).Encode(conf)
			filePath += ".json"
		default:
			err = fmt.Errorf("%s is not a valid configuration file format", initFormat)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(filePath, buf.Bytes(), 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}
