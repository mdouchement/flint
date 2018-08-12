package lint

/*
type Rule struct {
	Name     string
	Files    []string
	Style    string
	Pattern  string
	Severity Severity
}
*/
// Rule defines an abstract rule interaface
type Rule interface {
	Name() string
	//Category() string
	Apply(File) []Issue
}
