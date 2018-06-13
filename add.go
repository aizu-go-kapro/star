package main

import "github.com/mitchellh/cli"

type AddCommand struct {
}

func (a *AddCommand) Help() string {
	return "Usage: star add <url> <name>"
}

func (a *AddCommand) Run(args []string) int {
	panic("not implemented")
}

func (a *AddCommand) Synopsis() string {
	return "add a bookmark"
}

func newAddCommand() (cli.Command, error) {
	return &AddCommand{}, nil
}
