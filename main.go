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

var verbose bool

func main() {
	var (
		dbPath string
	)
	// TODO: fix help text
	flag.StringVar(&dbPath, "path", "", "JSON database path")
	flag.BoolVar(&verbose, "V", false, "verbose mode")
	flag.Parse()

	InitDB(dbPath)
	InitUI(nil)

	c := cli.NewCLI(appName, version)
	c.Args = flag.Args()
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
