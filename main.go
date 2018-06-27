package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

const (
	appName = "star"
	version = "0.1.0"
)

func main() {
	Init()

	c := cli.NewCLI(appName, version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"open":   newOpenCommand,
		"add":    newAddCommand,
		"delete": newDeleteCommand,
		"list":   newListCommand,
		"update": newUpdateCommand,
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(exitStatus)
}
