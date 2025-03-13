/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/rezaAmiri123/ormus/cli/cmd"
	_ "github.com/rezaAmiri123/ormus/cli/cmd/config"
	_ "github.com/rezaAmiri123/ormus/cli/cmd/destination"
	_ "github.com/rezaAmiri123/ormus/cli/cmd/project"
	_ "github.com/rezaAmiri123/ormus/cli/cmd/source"
	_ "github.com/rezaAmiri123/ormus/cli/cmd/user"
)

func main() {
	cmd.Execute()
}
