package main

import "github.com/mitchellh/cli"

type ListCommand struct{}

func (l *ListCommand) Help() string {
	return "Usage: star list"
}

func (l *ListCommand) Run(args []string) int {
	panic("not implemented")
}

func (l *ListCommand) Synopsis() string {
	return "list ups all bookmarks"
}

func newListCommand() (cli.Command, error) {
	return &ListCommand{}, nil
}
