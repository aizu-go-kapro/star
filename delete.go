package main

import "github.com/mitchellh/cli"

type DeleteCommand struct {
}

func (d *DeleteCommand) Help() string {
	panic("not implemented")
}

func (d *DeleteCommand) Run(args []string) int {
	panic("not implemented")
}

func (d *DeleteCommand) Synopsis() string {
	panic("not implemented")
}

func newDeleteCommand() (cli.Command, error) {
	return &DeleteCommand{}, nil
}
