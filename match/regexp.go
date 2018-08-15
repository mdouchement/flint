package match

import (
	"fmt"
	"regexp"
)

type Regexp struct {
	patterns []*regexp.Regexp
}

func (reg Regexp) MatchString(str string) bool {
	for _, pattern := range reg.patterns {
		if pattern.MatchString(str) {
			return true
		}
	}
	return false
}

func (reg Regexp) ToStringSlice() []string {
	ret := make([]string, len(reg.patterns))

	for i, reg := range reg.patterns {
		ret[i] = reg.String()
	}

	return ret
}

func NewRegexp(patterns []string) (Regexp, error) {
	ret := Regexp{patterns: make([]*regexp.Regexp, len(patterns))}

	for i, pattern := range patterns {
		reg, err := regexp.Compile(pattern)
		if err != nil {
			return ret, fmt.Errorf("invalid regexp pattern: %s: %s", pattern, err.Error())
		}
		ret.patterns[i] = reg
	}
	return ret, nil
}
