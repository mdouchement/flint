package lint

type Rule struct {
	Name     string
	Files    []string
	Style    string
	Pattern  string
	Severity Severity
}
