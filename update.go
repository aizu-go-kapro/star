package main

import "github.com/mitchellh/cli"

type UpdateCommand struct{}

func (u *UpdateCommand) Help() string {
	panic("not implemented")
}

func (u *UpdateCommand) Run(args []string) int {
	panic("not implemented")
}

func (u *UpdateCommand) Synopsis() string {
	panic("not implemented")
}

func newUpdateCommand() (cli.Command, error) {
	return &UpdateCommand{}, nil
}
