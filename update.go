package main

import (
	"context"
	"fmt"

	"github.com/mitchellh/cli"
)

type UpdateCommand struct {
	ui UI
}

func (*UpdateCommand) Help() string {
	return "Usage: star update <url> <name>"
}

func (u *UpdateCommand) Run(args []string) int {
	if len(args) != 2 {
		u.ui.ErrPrintln(u.Help())
		return 1
	}

	if err := repo.Bookmark.Update(context.Background(), &Bookmark{Name: args[1], URL: args[0]}); err != nil {
		u.ui.ErrPrintln(fmt.Sprintf("failed to update the bookmark: %s", err))
		return 1
	}

	return 0
}

func (*UpdateCommand) Synopsis() string {
	return "update a bookmark"
}

func newUpdateCommand() (cli.Command, error) {
	return &UpdateCommand{ui: ui}, nil
}
