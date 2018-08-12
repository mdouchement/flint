package lint

import (
	"errors"
)

const (
	// SeverityOff declare ignored issues.
	SeverityOff = "off"
	// SeverityWarning declares issues of type warning.
	SeverityWarning = "warning"
	// SeverityError declares issues  of type error.
	SeverityError = "error"
)

// Severity is the type for the failure types.
type Severity string

// Issue defines a struct for a linting issue.
type Issue struct {
	// the path to the file or directory
	Path            string
	ViolatedStyle   string
	ViolatedPattern string
	Severity        Severity
	// the rule's number
	Rule uint64
}

// ValidateSeverity check if the given severity is valdi
func ValidateSeverity(severity Severity) error {
	if severity != "" && severity != SeverityOff && severity != SeverityWarning && severity != SeverityError {
		return errors.New(string(severity) + " is not a valid severity")
	}
	return nil
}
