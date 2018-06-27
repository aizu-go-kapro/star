package main

import "github.com/mitchellh/cli"

type OpenCommand struct {
}

func (o *OpenCommand) Help() string {
	return "Usage: star open <name>"
}

func (o *OpenCommand) Run(args []string) int {
	panic("not implemented")
}

func (o *OpenCommand) Synopsis() string {
	return "open the url of a bookmark you selected"
}

func newOpenCommand() (cli.Command, error) {
	return &OpenCommand{}, nil
}
