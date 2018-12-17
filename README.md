# Flint

[![GoDoc](https://godoc.org/github.com/bloom42/flint?status.svg)](https://godoc.org/github.com/bloom42/flint)
[![GitHub release](https://img.shields.io/github/release/bloom42/flint.svg)](https://github.com/bloom42/flint/releases/latest)
[![Build Status](https://travis-ci.org/bloom42/flint.svg?branch=master)](https://travis-ci.org/bloom42/flint)

You think there is no place in code reviews to discuss about files naming conventions ? Impose consistent files and directories naming rules with **flint**: the filesystem linter.

1. [Installation](#installation)
2. [Docker image](#docker-image)
3. [Usage](#usage)
4. [Configuration](#configuration)
5. [Available rules](#available-rules)
6. [Available formatters](#available-formatters)
7. [Roadmap](#roadmap)

-------------------

## Installation

### Using go (nightly)
```bash
$ go get -u github.com/bloom42/flint
```

### Binary releases
[https://github.com/bloom42/flint/releases/latest](https://github.com/bloom42/flint/releases/latest)




## Docker image

[bloom42/flint](https://hub.docker.com/r/bloom42/flint)




## Usage

Go to your project's root directory then
```bash
$ flint init # create a configuration file with default configuration
$ flint
# or cd my_directory && flint to lint only current directory and subfiles
```




## Configuration

Configuration is stored in a `.flint.sane` file at the root of your project (repo).

When runned, flint will recursively search upward for a configuraiton file, starting in the current
directory.

```bash
$ cat .flint.sane
```

```sane
description = "This is a configuration file for flint, the filesystem linter. More information here: https://github.com/z0mbie42/flint"
default_severity = "warning" # valid values are [off, warning, error]
error_exit_code = 1
warning_exit_code = 0
match_format = "blob" # match format for ignore_directories and ignore_files, valid values are [blob, regexp]


# you can ignore files and directories using glob or regexp syntax according to the configuration above
ignore_files = [".*", "vendor", "Gopkg.toml", "Gopkg.lock", "README.md", "LICENSE"]
ignore_directories = [".*", "vendor"]


# define used rules
rules = {
  "dir/no_dot" = {}
  "file/lower_case_ext" = {}
  "file/no_multi_ext" = {}
  no_empty_name = {}
  no_leading_underscores = {}
  no_trailing_underscores = {}
  no_whitespaces = {}
  snake_case = {
    severity = "error"
  }
}
```



## Available rules

| Name                  | Description                                                              |
| --------------------- | :----------------------------------------------------------------------- |
| `dir/no_dot`          | Disallows the usage of `.` in directory names.                           |
| `file/lower_case_ext` | Files names should be lower cased.                                       |
| `file/no_multi_ext`   | Disallows usage of multiple extensions in file names (eg. main.text.js). |
| `no_empty_name`       | Disallows usage of names which only contains whitespaces.                |
| `no_leading_underscores` | Disallows usage of leading underscores in names.                      |
| `no_trailing_underscores` | Disallows usage of trailing underscores in names.                    |
| `no_whitespaces`      | Disallows usage of whitespaces in names.                                 |
| `snake_case`          | Force names to respect the `snake case` convention (only lower cased alphanum and underscores). |




## Available formatters


### Console
![Console formatter](_docs/formatter_console.png)



### Basic
![Basic formatter](_docs/formatter_basic.png)



### Json
![Json formatter](_docs/formatter_json.png)



### Ndjson
![Ndson formatter](_docs/formatter_ndjson.png)





## Roadmap

See [https://github.com/bloom42/flint/projects/1](https://github.com/bloom42/flint/projects/1)
