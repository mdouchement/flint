package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/astrocorp42/flint/formatter"
	"github.com/astrocorp42/flint/lint"
	"github.com/astrocorp42/flint/match"
	"github.com/astrocorp42/flint/rule"
	"github.com/astrocorp42/flint/rule/dir"
	"github.com/astrocorp42/flint/rule/file"
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

var AllFormatters = []lint.Formatter{
	formatter.Console{},
	formatter.Basic{},
	formatter.JSON{},
	formatter.NDJSON{},
}

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
	config := lint.ConfigFile{}
	var ret lint.Config
	ext := filepath.Ext(configFilePath)
	var err error

	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return ret, err
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
		return ret, err
	}

	return ConfigFileToConfig(config)
}

func Get() (lint.Config, error) {
	var err error
	var config lint.Config

	configFilePath := FindConfigFile()

	if configFilePath == "" {
		return config, errors.New(".flint(.toml|json) configuraiton file not found. Please run \"flint init\"")
	}

	config, err = parseConfig(configFilePath)
	if err != nil {
		return config, err
	}

	config.BaseDir = filepath.Dir(configFilePath)
	config.WorkingDir, err = os.Getwd()
	return config, err
}

func Default() lint.ConfigFile {
	config := lint.ConfigFile{Rules: lint.RulesConfig{}}

	config.Description = "This is a configuration file for flint, the filesystem linter. More " +
		"information here: https://github.com/astrocorp42/flint"
	config.DefaultSeverity = "warning"
	config.WarningExitCode = 0
	config.ErrorExitCode = 1
	config.MatchFormat = "blob"
	config.IgnoreFiles = []string{".*", "LICENSE*", "README*", "Gopkg.toml", "Gopkg.lock", "package.json", "Dockerfile", "Makefile", "package-lock.json"}
	config.IgnoreDirectories = []string{".*"}

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
	var err error

	ret.DefaultSeverity = configFile.DefaultSeverity
	ret.ErrorExitCode = configFile.ErrorExitCode
	ret.WarningExitCode = configFile.WarningExitCode

	if configFile.MatchFormat == "" {
		ret.MatchFormat = "blob"
	} else {
		ret.MatchFormat = configFile.MatchFormat
	}

	ret.Rules = make(lint.Rules, len(configFile.Rules))
	i := 0
	for name := range configFile.Rules {
		rule, ok := findRule(AllRules, name)
		if !ok {
			return ret, fmt.Errorf("cannot find rule: %s", name)
		}
		ret.Rules[i] = rule
		i += 1
	}

	ret.RulesConfig = configFile.Rules

	switch ret.MatchFormat {
	case "blob":
		ret.IgnoreFiles, err = match.NewBlob(configFile.IgnoreFiles)
		if err != nil {
			break
		}
		ret.IgnoreDirectories, err = match.NewBlob(configFile.IgnoreDirectories)
	case "regexp":
		ret.IgnoreFiles, err = match.NewRegexp(configFile.IgnoreFiles)
		if err != nil {
			break
		}
		ret.IgnoreDirectories, err = match.NewRegexp(configFile.IgnoreDirectories)
	default:
		return ret, fmt.Errorf("match_format is not valid: %s, accepted values are [blob, regexp]", ret.MatchFormat)
	}

	return ret, nil
}
