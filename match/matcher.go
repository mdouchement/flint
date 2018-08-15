package match

type Matcher interface {
	MatchString(string) bool
	ToStringSlice() []string
}
