package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/astrocorp42/flint/version"
	"github.com/spf13/cobra"
)

var versionFormat string

func init() {
	RootCmd.AddCommand(VersionCmd)
	VersionCmd.PersistentFlags().StringVarP(&versionFormat, "format", "f", "console", "Output format. Valid values are [console, json]")
}

type versionJSON struct {
	Version      string `json:"version"`
	GitCommit    string `json:"git_commit"`
	UTCBuildTime string `json:"utc_build_time"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	GoVersion    string `json:"go_version"`
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version and build information",
	Long:  "Display the version and build information",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var output []byte

		switch versionFormat {
		case "console":
			fmt.Printf("Version        : %s\n", version.Version)
			fmt.Printf("Git commit     : %s\n", version.GitCommit)
			fmt.Printf("UTC build time : %s\n", version.UTCBuildTime)
			fmt.Printf("OS             : %s\n", version.OS)
			fmt.Printf("Architecture   : %s\n", version.Arch)
			fmt.Printf("Go version     : %s\n", version.GoVersion)
		case "json":
			jsonVersion := versionJSON{
				Version:      version.Version,
				GitCommit:    version.GitCommit,
				UTCBuildTime: version.UTCBuildTime,
				OS:           version.OS,
				Architecture: version.Arch,
				GoVersion:    version.GoVersion,
			}
			output, err = json.Marshal(&jsonVersion)
			if err == nil {
				fmt.Println(string(output))
			}
		default:
			err = fmt.Errorf("Error: %s is not a valid output format", versionFormat)

		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(3)
		}

	},
}
