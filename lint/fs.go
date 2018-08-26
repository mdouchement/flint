package lint

import (
	"github.com/z0mbie42/fswalk"
)

type File struct {
	fswalk.File
	RelativePath string  `json:"relative_path"`
	Ext          string  `json:"ext"`
	Issues       []Issue `json:"issues"`
}
