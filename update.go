package main

import "github.com/mitchellh/cli"

type UpdateCommand struct {
	ui UI
}

func (*UpdateCommand) Help() string {
	return "Usage: star update <url> <name>"
}

func (u *UpdateCommand) Run(args []string) int {
	panic("not implemented")
}

func (*UpdateCommand) Synopsis() string {
	return "update a bookmark"
}

func newUpdateCommand() (cli.Command, error) {
	return &UpdateCommand{ui: ui}, nil
}
