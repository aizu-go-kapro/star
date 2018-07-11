package main

import "github.com/mitchellh/cli"

type DeleteCommand struct {
	ui UI
}

func (d *DeleteCommand) Help() string {
	return "Usage: star delete <name>"
}

func (d *DeleteCommand) Run(args []string) int {
	panic("not implemented")
}

func (d *DeleteCommand) Synopsis() string {
	return "delete a bookmark"
}

func newDeleteCommand() (cli.Command, error) {
	return &DeleteCommand{ui: ui}, nil
}
