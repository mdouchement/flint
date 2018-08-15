package lint

type Matcher interface {
	MatchString(string) bool
	ToStringSlice() []string
}
