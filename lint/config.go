package lint

// Arguments is type used for the arguments of a rule.
type Arguments = []interface{}

// RuleConfig is type used for the rule configuration.
type RuleConfig struct {
	Arguments Arguments `toml:"arguments,omitempty" json:"arguments,omitempty"`
	Severity  Severity  `toml:"severity,omitempty" json:"severity,omitempty"`
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
	BasePath    string   `toml:"-" json:"-"`
	WorkingDir  string   `toml:"-" json:"-"`
	Format      string   `toml:"format" json:"format"`     // default output format
	Severity    Severity `toml:"severity" json:"severity"` // default severity
	ErrorCode   int      `toml:"error_code" json:"error_code"`
	WarningCode int      `toml:"warning_code" json:"warning_code"`

	Comment string `toml:"comment" json:"comment"`
	//Files             []string    `toml:"files" json:"files"`
	//Directories       []string    `toml:"directories" json:"directories"`
	Rules             RulesConfig `toml:"rules" json:"rules"`
	IgnoreFiles       []string    `toml:"ignore_files" json:"ignore_files"`
	IgnoreDirectories []string    `toml:"ignore_directories" json:"ignore_directories"`
}
