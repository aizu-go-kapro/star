package main

import (
	"context"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/mitchellh/cli"
	"github.com/pkg/errors"
)

type DeleteCommand struct {
	ui UI
}

func (d *DeleteCommand) Help() string {
	return "Usage: star delete <name>"
}

func (d *DeleteCommand) Run(args []string) int {
	if len(args) == 0 {
		d.ui.ErrPrintln(d.Help())
		return 1
	}

	var result error
	for _, name := range args {
		if err := repo.Bookmark.Delete(context.Background(), &Bookmark{Name: name}); err != nil {
			result = multierror.Append(result, errors.Wrapf(err, "no such bookmark: %s", name))
		}
	}
	if result != nil {
		d.ui.ErrPrintln(fmt.Sprintf("failed to delete some bookmarks: %s", result))
		return 1
	}
	return 0
}

func (d *DeleteCommand) Synopsis() string {
	return "delete a bookmark"
}

func newDeleteCommand() (cli.Command, error) {
	return &DeleteCommand{ui: ui}, nil
}
