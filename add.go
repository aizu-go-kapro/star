package main

import (
	"context"
	"net/url"
	"time"

	"github.com/mitchellh/cli"
)

type AddCommand struct {
	ui UI
}

func (a *AddCommand) Help() string {
	return "Usage: star add <url> <name>"
}

func (a *AddCommand) Run(args []string) int {
	if len(args) != 2 {
		a.ui.Println(a.Help())
		return 1
	}

	if _, err := url.ParseRequestURI(args[0]); err != nil {
		a.ui.ErrPrintln(err)
		return 1
	}

	bookmark := &Bookmark{Name: args[1], URL: args[0], CreatedAt: time.Now()}
	if err := repo.Bookmark.Add(context.Background(), bookmark); err != nil {
		a.ui.Println(err)
		return 1
	}

	return 0
}

func (a *AddCommand) Synopsis() string {
	return "add a bookmark"
}

func newAddCommand() (cli.Command, error) {
	return &AddCommand{ui: ui}, nil
}
