package main

import (
	flags "github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(nil, flags.Default)
	parser.AddCommand("release", "Create a new release", "Create a new release from the latest commit.", &Release)
	parser.Parse()
}
