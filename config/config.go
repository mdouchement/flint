package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/z0mbie42/flint/lint"
	"github.com/z0mbie42/flint/rule"
	"github.com/z0mbie42/flint/rule/dir"
	"github.com/z0mbie42/flint/rule/file"
)

const DefaultConfigurationFileName = ".flint"

var DefaultRules = lint.Rules{
	rule.NoLeadingUnderscores{},
	rule.NoTrailingUnderscores{},
	rule.NoEmptyName{},
	rule.NoWhitespaces{},
	rule.SnakeCase{},
	file.LowerCaseExt{},
	file.NoMultiExt{},
	dir.NoDot{},
}

var AllRules = lint.Rules{
	rule.NoLeadingUnderscores{},
	rule.NoTrailingUnderscores{},
	rule.NoEmptyName{},
	rule.NoWhitespaces{},
	rule.SnakeCase{},
	file.LowerCaseExt{},
	file.NoMultiExt{},
	dir.NoDot{},
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

func parseConfig(configFilePath string) (lint.ConfigFile, error) {
	config := lint.ConfigFile{}
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
		config.IgnoreFiles = []*regexp.Regexp{}
	}
	if config.IgnoreDirectories == nil {
		config.IgnoreDirectories = []*regexp.Regexp{}
	}
	return nil
}

func Get() (lint.Config, error) {
	var err error
	configFile := lint.ConfigFile{Rules: lint.RulesConfig{}}
	var config lint.Config

	configFilePath := FindConfigFile()

	if configFilePath == "" {
		return config, errors.New(".flint(.toml|json) configuraiton file not found. Please run \"flint init\"")
	}

	configFile, err = parseConfig(configFilePath)
	if err != nil {
		return config, err
	}

	config, err = ConfigFileToConfig(configFile)
	if err != nil {
		return config, err
	}

	if err = normalizeConfig(&config); err != nil {
		return config, err
	}

	config.BaseDir = filepath.Dir(configFilePath)
	config.WorkingDir, err = os.Getwd()
	if err != nil {
		return config, err
	}

	return config, nil
}

func Default() lint.ConfigFile {
	config := lint.ConfigFile{Rules: lint.RulesConfig{}}

	config.Comment = "This is a configuration file for flint, the filesystem linter. More " +
		"information here: https://github.com/z0mbie42/flint"
	config.Format = "default"
	config.Severity = "warning"
	config.WarningCode = 0
	config.ErrorCode = 1
	config.IgnoreFiles = []string{"(^|/)\\..*"}
	config.IgnoreDirectories = []string{"(^|/)\\..*"}

	for _, rule := range DefaultRules {
		config.Rules[rule.Name()] = lint.RuleConfig{}
	}

	return config
}

func findRule(arr []lint.Rule, name string) (lint.Rule, bool) {
	for _, rule := range arr {
		if rule.Name() == name {
			return rule, true
		}
	}
	return nil, false
}

func ConfigFileToConfig(configFile lint.ConfigFile) (lint.Config, error) {
	ret := lint.Config{}

	ret.Format = configFile.Format
	ret.Severity = configFile.Severity
	ret.ErrorCode = configFile.ErrorCode
	ret.WarningCode = configFile.WarningCode

	ret.Rules = lint.Rules{}
	for name := range configFile.Rules {
		rule, ok := findRule(AllRules, name)
		if !ok {
			return ret, fmt.Errorf("cannot find rule: %s", name)
		}
		ret.Rules = append(ret.Rules, rule)
	}

	ret.RulesConfig = configFile.Rules

	ret.IgnoreFiles = []*regexp.Regexp{}
	for _, regex := range configFile.IgnoreFiles {
		reg, err := regexp.Compile(regex)
		if err != nil {
			return ret, fmt.Errorf("invalid regexp pattern: %s: %s", regex, err.Error())
		}
		ret.IgnoreFiles = append(ret.IgnoreFiles, reg)
	}

	ret.IgnoreDirectories = []*regexp.Regexp{}
	for _, regex := range configFile.IgnoreDirectories {
		reg, err := regexp.Compile(regex)
		if err != nil {
			return ret, fmt.Errorf("invalid regexp pattern: %s: %s", regex, err.Error())
		}
		ret.IgnoreDirectories = append(ret.IgnoreDirectories, reg)
	}

	return ret, nil
}
