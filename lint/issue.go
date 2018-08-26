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
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

// Severity is the type for the failure types.
type Severity string

// Issue defines a struct for a linting issue.
type Issue struct {
	//Severity Severity
	Rule     string        `json:"rule"`
	Message  string        `json:"message"`
	Position IssuePosition `json:"position"`
	Severity Severity      `json:"severity"`
}

// ValidateSeverity check if the given severity is valdi
func ValidateSeverity(severity Severity) error {
	if severity != "" && severity != SeverityOff && severity != SeverityWarning && severity != SeverityError {
		return errors.New(string(severity) + " is not a valid severity")
	}
	return nil
}
