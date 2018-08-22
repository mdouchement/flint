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

type IssuePosition struct {
	Start uint64
	End   uint64
}

// Severity is the type for the failure types.
type Severity string

// Issue defines a struct for a linting issue.
type Issue struct {
	//Severity Severity
	File     File
	RuleName string
	Message  string
	Position IssuePosition
}

// ValidateSeverity check if the given severity is valdi
func ValidateSeverity(severity Severity) error {
	if severity != "" && severity != SeverityOff && severity != SeverityWarning && severity != SeverityError {
		return errors.New(string(severity) + " is not a valid severity")
	}
	return nil
}
