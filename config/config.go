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
	"github.com/z0mbie42/flint/rule"
)

const DefaultConfigurationFileName = ".flint"

var DefaultRules = []lint.Rule{
	rule.NoLeadingUnderscores{},
	rule.NoTrailingUnderscores{},
	rule.NoDirDot{},
	rule.NoEmptyName{},
	rule.NoMultiExtensions{},
	rule.NoWhitespaces{},
	rule.SnakeCase{},
}

var AllRules = []lint.Rule{
	rule.NoLeadingUnderscores{},
	rule.NoTrailingUnderscores{},
	rule.NoDirDot{},
	rule.NoEmptyName{},
	rule.NoMultiExtensions{},
	rule.NoWhitespaces{},
	rule.SnakeCase{},
}

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

func parseConfig(configFilePath string) (lint.Config, error) {
	config := lint.Config{}
	ext := filepath.Ext(configFilePath)
	var err error

	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return config, err
	}

	switch ext {
	case ".toml":
		_, err = toml.Decode(string(file), &config)
	case ".json":
		err = json.Unmarshal(file, &config)
	default:
		err = errors.New(ext + " is not a configuration file extension")
	}
	if err != nil {
		return config, err
	}

	return config, nil
}

func normalizeConfig(config *lint.Config) error {
	if config.IgnoreFiles == nil {
		config.IgnoreFiles = []string{}
	}
	if config.IgnoreDirectories == nil {
		config.IgnoreDirectories = []string{}
	}
	return nil
}

func Get() (lint.Config, error) {
	var err error
	config := lint.Config{Rules: map[string]lint.RuleConfig{}}

	configFilePath := FindConfigFile()

	if configFilePath == "" {
		return config, errors.New(".flint(.toml|json) configuraiton file not found. Please run \"flint init\"")
	}

	config, err = parseConfig(configFilePath)
	if err != nil {
		return config, err
	}

	if err = normalizeConfig(&config); err != nil {
		return config, err
	}

	config.BasePath = filepath.Dir(configFilePath)

	return config, nil
}

func Default() lint.Config {
	config := lint.Config{Rules: map[string]lint.RuleConfig{}}

	config.Comment = "This is a configuration file for flint, the filesystem linter. More " +
		"information here: https://github.com/z0mbie42/flint"
	config.Format = "default"
	config.Severity = "warning"
	config.WarningCode = 0
	config.ErrorCode = 1
	//config.Directories = []string{"**"}
	//config.Files = []string{"**"}
	config.IgnoreFiles = []string{"**/.*", "**/.*/**/.*", "**/.*/**/*"}
	config.IgnoreDirectories = []string{"**/.*", ".**/*"}

	for _, rule := range DefaultRules {
		config.Rules[rule.Name()] = lint.RuleConfig{}
	}

	return config
}

func GetLoadedRules(config lint.Config) ([]lint.Rule, error) {
	rulesMap := map[string]lint.Rule{}
	for _, r := range AllRules {
		rulesMap[r.Name()] = r
	}

	lintingRules := []lint.Rule{}
	for name := range config.Rules {
		rule, ok := rulesMap[name]
		if !ok {
			return nil, errors.New("cannot find rule: " + name)
		}
		lintingRules = append(lintingRules, rule)
	}

	return lintingRules, nil
}
