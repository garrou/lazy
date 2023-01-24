package main

import (
	_ "embed"
	"lazy/cmd"
	"lazy/lib"
)

//go:embed data/lazy.json
var filename string

func main() {
	lib.Filename = filename
	cmd.Execute()
}
