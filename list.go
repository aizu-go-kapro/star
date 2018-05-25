package main

import "github.com/mitchellh/cli"

type ListCommand struct{}

func (l *ListCommand) Help() string {
	panic("not implemented")
}

func (l *ListCommand) Run(args []string) int {
	panic("not implemented")
}

func (l *ListCommand) Synopsis() string {
	panic("not implemented")
}

func newListCommand() (cli.Command, error) {
	return &ListCommand{}, nil
}
