package main

import (
	"flag"
	"os"

	"github.com/mitchellh/cli"
)

const (
	appName = "star"
	version = "0.1.0"
)

func main() {
	var (
		dbPath string
	)
	flag.StringVar(&dbPath, "path", "", "JSON database path")
	flag.Parse()

	InitDB(dbPath)
	InitUI(nil)

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
		ui.ErrPrintln(err)
	}
	os.Exit(exitStatus)
}
