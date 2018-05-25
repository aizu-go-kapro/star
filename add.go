package main

import "github.com/mitchellh/cli"

type AddCommand struct {
}

func (a *AddCommand) Help() string {
	panic("not implemented")
}

func (a *AddCommand) Run(args []string) int {
	panic("not implemented")
}

func (a *AddCommand) Synopsis() string {
	panic("not implemented")
}

func newAddCommand() (cli.Command, error) {
	return &AddCommand{}, nil
}
