package lint

type File struct {
	Name  string
	Path  string
	Ext   string
	IsDir bool
}

func (file File) lint(config Config, loadedRules []Rule, errorsc <-chan error) []Issue {
	foundIssues := []Issue{}

	for _, currentRule := range loadedRules {
		foundIssues = append(foundIssues, currentRule.Apply(file)...)
	}

	return foundIssues
}
