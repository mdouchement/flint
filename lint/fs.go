package lint

type File struct {
	Name         string  `json:"name"`
	Path         string  `json:"path"`
	RelativePath string  `json:"relative_path"`
	Ext          string  `json:"ext"`
	IsDir        bool    `json:"is_dir"`
	Issues       []Issue `json:"issues"`
}
