package lint

// Arguments is type used for the arguments of a rule.
type Arguments = []interface{}

// RuleConfig is type used for the rule configuration.
type RuleConfig struct {
	Arguments Arguments `toml:"arguments,omitempty"`
	Severity  Severity  `toml:"severity,omitempty"`
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
	BasePath    string      `toml:"-"`
	Description string      `toml:"description"`
	Format      string      `toml:"format"`
	Severity    Severity    `toml:"severity"`
	Rules       RulesConfig `toml:"rules"`
	ErrorCode   int         `toml:"error_code"`
	WarningCode int         `toml:"warning_code"`
	//IgnoredFiles       []string     `toml:"ignored_files"`
	//IgnoredDirectories []string     `toml:"ignored_directories"`
	//DefinedRules DefinedRules `toml:"defined_rules"`
}
