# Flint

A fast and configurable filesystem (file and directory names) linter

[![GoDoc](https://godoc.org/github.com/z0mbie42/flint?status.svg)](https://godoc.org/github.com/z0mbie42/flint)
[![GitHub release](https://img.shields.io/github/release/z0mbie42/flint.svg)](https://github.com/z0mbie42/flint)

1. [Configuration](#configuration)
2. [Usage](#usage)
3. [Roadmap](#roadmap)

-------------------

## Configuration

Configuration is stored in the `.flint.(toml|json)` file.

When runned, flint will recursively search upward for a configuraiton file, starting in the current
directory.

```bash
$ cat .flint.toml
```


```toml
description = "this is a configuration file for flint (https://github.com/z0mbie42/flint"

[[rules]]
# rule for files
files = [
  "assets/js/**/*.js",
  "lib/{,*/}*.js",
  "routes/*."
]

# rule for directories
directories = [

]

# ignore specific files wichi already matched the above patterns
ignored_files = [ "**/node_modules/**" ] # default to []
ignore_directories = [] # default to []

# you can use a predefined style [snake, kebab, pascal, camel], it will check parts between dots eg for main.go: "main" and "go"
style = "snake"
# or a regex pattern (should take in account the dots)
pattern = "[a-z]+[a-z_]*[a-z]+\.[a-z]+"
# or both (should match both)

severity = "error" # default to error, valid values are: [off, warning, error]
```



## Usage

```bash
$ flint
```


## Roadmap

- [ ] Add the `fix` (with the `--plan` option) command to automagically fix issues.
