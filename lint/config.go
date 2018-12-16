package lint

import (
	"github.com/bloom42/flint/match"
)

// Arguments is type used for the arguments of a rule.
type Arguments = []interface{}

// RuleConfig is type used for the rule configuration.
type RuleConfig struct {
	Arguments Arguments `sane:"arguments,omitempty" json:"arguments,omitempty"`
	Severity  Severity  `sane:"severity,omitempty" json:"severity,omitempty"`
}

// RulesConfig defines the config for all rules.
type RulesConfig = map[string]RuleConfig

// DefinedRule is an user defined rule
type DefinedRule struct {
	Description string
	Pattern     string
	Type        string // file, directory or both
}

// DefinedRules defines the config for all the user defined rules
type DefinedRules = map[string]DefinedRule

// Config defines the config of the linter.
type Config struct {
	// the directory of the config file, relative to the execution of flint
	// Extends []strings // extend a set of rules
	// for json which does not have comments
	BaseDir         string
	WorkingDir      string
	DefaultSeverity Severity
	ErrorExitCode   int
	WarningExitCode int
	MatchFormat     string

	//Files             []string    `sane:"files" json:"files"`
	//Directories       []string    `sane:"directories" json:"directories"`
	Rules             Rules
	RulesConfig       RulesConfig
	IgnoreFiles       match.Matcher
	IgnoreDirectories match.Matcher
}

func (config Config) ToFile() ConfigFile {
	ret := ConfigFile{}

	ret.DefaultSeverity = config.DefaultSeverity
	ret.ErrorExitCode = config.ErrorExitCode
	ret.WarningExitCode = config.WarningExitCode

	ret.Rules = RulesConfig{}
	for name, config := range config.RulesConfig {
		ret.Rules[name] = config
	}

	ret.IgnoreFiles = config.IgnoreFiles.ToStringSlice()
	ret.IgnoreDirectories = config.IgnoreDirectories.ToStringSlice()

	return ret
}

type ConfigFile struct {
	Description     string   `sane:"description" json:"description"`
	DefaultSeverity Severity `sane:"default_severity" json:"default_severity"`
	ErrorExitCode   int      `sane:"error_exit_code" json:"error_exit_code"`
	WarningExitCode int      `sane:"warning_exit_code" json:"warning_exit_code"`
	MatchFormat     string   `sane:"match_format" json:"match_format"`

	Rules             RulesConfig `sane:"rules" json:"rules"`
	IgnoreFiles       []string    `sane:"ignore_files" json:"ignore_files"`
	IgnoreDirectories []string    `sane:"ignore_directories" json:"ignore_directories"`
}
