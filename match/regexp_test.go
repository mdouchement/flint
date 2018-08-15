package match

import (
	"testing"
)

type test struct {
	patterns []string
	str      string
	expected bool
}

func TestRegexpMatchString(t *testing.T) {
	tests := []test{
		{[]string{".*"}, "lol", true},
	}

	for i, test := range tests {
		matcher, err := NewRegexp(test.patterns)
		if err != nil {
			t.Errorf("invalid regexp: %v: %v", test.patterns, err)
		}
		res := matcher.MatchString(tests[i].str)
		if res != tests[i].expected {
			t.Errorf("incorrect regexp MatchString result: got %v, expected: %v, str: %v, patterns: %v", res, tests[i].expected, tests[i].str, test.patterns)
		}
	}
}
