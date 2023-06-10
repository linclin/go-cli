package main

import (
	"go-cli/cmd"
	_ "go-cli/cmd/todo"
	_ "go-cli/cmd/version"
)

func init() {
}

func main() {
	cmd.Execute()
}
