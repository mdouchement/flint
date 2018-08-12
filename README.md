# Flint

Fast and configurable filesystem (file and directory names) linter

[![GoDoc](https://godoc.org/github.com/z0mbie42/flint?status.svg)](https://godoc.org/github.com/z0mbie42/flint)
[![GitHub release](https://img.shields.io/github/release/z0mbie42/flint.svg)](https://github.com/z0mbie42/flint)

1. [Configuration](#configuration)
2. [Usage](#usage)
3. [Roadmap](#roadmap)

-------------------

## Configuration

Configuration is stored in a `.flint.(toml|json)` file at the root of your project (repo).

When runned, flint will recursively search upward for a configuraiton file, starting in the current
directory.

```bash
$ cat .flint.toml
```


```toml
description = "this is a configuration file for flint (https://github.com/z0mbie42/flint)"
formatter = "default" # valid values are [default]

# Sets the error code for failures with severity "error"
error_code = 1 # default 1

# Sets the error code for failures with severity "warning"
warning_code = 0 # default 0

# Sets the default severity to "warning"
default_severity = "warning"

[rules.js_files]
    description = "" # optionnal rule description (eg: for json which does not have comments)
  files = [
    "assets/js/**/*.js",
    "lib/{,*/}*.js",
  ]

  # rule for directories
  directories = []

  # ignore specific files wichi already matched the above patterns
  ignored_files = [ "**/node_modules/**" ] # default to []
  ignore_directories = [] # default to []

  # you can use a predefined style [snake, kebab, pascal, camel], it will check parts between dots eg for main.go: "main" and "go"
  style = "snake"
  # or a regex pattern (should take in account the dots)
  pattern = "[a-z]+[a-z_]*[a-z]+\.[a-z]+"
  # or both (AND)

  severity = "error" # default to error [off, warning, error]
```



## Usage

```bash
$ flint init # create a configuration file with default configuration
$ flint
# or cd my_directory && flint . to lint only current directory and subfiles
```


## Roadmap

- [ ] Add the `fix` (with the `--plan` option) command to automagically fix issues.
