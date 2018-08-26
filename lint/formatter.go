package lint

// FormatterMetadata configuration of a formatter
type FormatterMetadata struct {
	Name        string
	Description string
	Sample      string
}

// Formatter defines an interface for issues formatters
type Formatter interface {
	Format(<-chan File) (<-chan string, <-chan error)
	Name() string
}
