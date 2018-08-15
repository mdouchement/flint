package match

import (
	"github.com/zealic/xignore"
)

type Blob struct {
	patterns []string
}

func (reg Blob) MatchString(str string) bool {
	ret, _ := xignore.IsMatch(str, reg.patterns)
	return ret
}

func (reg Blob) ToStringSlice() []string {
	ret := make([]string, len(reg.patterns))
	copy(ret, reg.patterns)
	return ret
}

func NewBlob(patterns []string) (Blob, error) {
	ret := Blob{patterns: make([]string, len(patterns))}
	copy(ret.patterns, patterns)
	return ret, nil
}
