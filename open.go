package main

import "github.com/mitchellh/cli"

type OpenCommand struct {
}

func (o *OpenCommand) Help() string {
	panic("not implemented")
}

func (o *OpenCommand) Run(args []string) int {
	panic("not implemented")
}

func (o *OpenCommand) Synopsis() string {
	panic("not implemented")
}

func newOpenCommand() (cli.Command, error) {
	return &OpenCommand{}, nil
}
