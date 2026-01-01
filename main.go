package main

import (
	"go-cli/cmd"
	_ "go-cli/cmd/todo"
	_ "go-cli/cmd/version"
	"go-cli/initialize"
)

func init() {
	initialize.SlogInit("info")
}

func main() {
	cmd.Execute()
}
