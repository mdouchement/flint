package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/z0mbie42/flint/lint"
)

const DefaultConfigurationFileName = ".flint"

var DefaultRules = []lint.Rule{}

var AllRules = []lint.Rule{}

var AllFormatters = []lint.Formatter{}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func FindConfigFile() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	directory := wd
	configFilePath := path.Join(directory, DefaultConfigurationFileName)

	for {
		directory = path.Clean(directory)
		if directory == "/" || directory == "." {
			break
		}

		configFilePath = path.Join(directory, DefaultConfigurationFileName)
		if FileExists(configFilePath + ".toml") {
			return configFilePath + ".toml"
		} else if FileExists(configFilePath + ".json") {
			return configFilePath + ".json"
		}

		directory = path.Dir(directory)
	}

	return ""
}

func parseConfig(configFilePath string) (*lint.Config, error) {
	config := &lint.Config{}
	ext := filepath.Ext(configFilePath)
	var err error

	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	switch ext {
	case ".toml":
		_, err = toml.Decode(string(file), config)
	case ".json":
		err = json.Unmarshal(file, config)
	default:
		err = errors.New(ext + " is not a configuration file extension")
	}
	if err != nil {
		return nil, err
	}

	return config, nil
}

func normalizeConfig(config *lint.Config) error {
	return nil
}

func Get() (*lint.Config, error) {
	var err error

	configFilePath := FindConfigFile()

	if configFilePath == "" {
		return nil, errors.New(".flint(.toml|json) configuraiton file not found. Please run \"flint init\"")
	}

	config, err := parseConfig(configFilePath)
	if err != nil {
		return nil, err
	}

	if err = normalizeConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func Default() *lint.Config {
	config := &lint.Config{}
	config.Description = "This is a configuration file for flint, the filesystem linter. More " +
		"information at https://github.com/z0mbie42/flint"
	config.Format = "default"
	config.Severity = "warning"
	config.WarningCode = 0
	config.ErrorCode = 1
	return config
}
