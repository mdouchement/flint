# Flint

Fast and configurable filesystem (file and directory names) linter

[![GoDoc](https://godoc.org/github.com/z0mbie42/flint?status.svg)](https://godoc.org/github.com/z0mbie42/flint)
[![GitHub release](https://img.shields.io/github/release/z0mbie42/flint.svg)](https://github.com/z0mbie42/flint)

1. [Installation](#installation)
2. [Usage](#usage)
3. [Configuration](#configuration)
4. [Roadmap](#roadmap)

-------------------

## Installation

```bash
$ go get -u github.com/z0mbie42/flint
```



## Usage

```bash
$ flint init # create a configuration file with default configuration
$ flint
# or cd my_directory && flint to lint only current directory and subfiles
```



## Configuration

Configuration is stored in a `.flint.(toml|json)` file at the root of your project (repo).

When runned, flint will recursively search upward for a configuraiton file, starting in the current
directory.

```bash
$ cat .flint.toml
```

```toml
# as json does not allow comments, you can use "comment" field everywhere
comment = "This is a configuration file for flint, the filesystem linter. More information here: https://github.com/z0mbie42/flint"
format = "default" # valid values are [default]
severity = "warning" # valid values are [off, warning, error]
error_code = 1
warning_code = 0
match_format = "blob" # match format for ignore_directories and ignore_files, valid values are [blob, regexp]


# you can ignore files and directories using glob or regexp syntax according to the configuration above
ignore_files = [".*", "vendor", "Gopkg.toml", "Gopkg.lock", "README.md", "LICENSE"]
ignore_directories = [".*", "vendor"]


# define used rules
[rules]
  [rules."dir/no_dot"]
  [rules."file/lower_case_ext"]
  [rules."file/no_multi_ext"]
  [rules.no_empty_name]
  [rules.no_leading_underscores]
  [rules.no_trailing_underscores]
  [rules.no_whitespaces]
  [rules.snake_case]
```



## Roadmap

- [ ] Add the `fix` command to automagically fix issues (with the `--plan` option).
