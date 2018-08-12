package main

type Config struct {
	Description string
	Rules       []Rule
}

type Rule struct {
	Files             []string
	Directories       []string
	IgnoreFiles       []string
	IgnoreDirectories []string
	Style             string
	Pattern           string
	Severity          string
}
