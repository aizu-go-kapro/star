package main

import (
	"context"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/mitchellh/cli"
	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
)

type OpenCommand struct {
	ui UI
}

func (o *OpenCommand) Help() string {
	return "Usage: star open <name>"
}

func (o *OpenCommand) Run(args []string) int {
	if len(args) == 0 {
		o.ui.ErrPrintln(o.Help())
		return 1
	}

	var result error
	for _, name := range args {
		b, err := repo.Bookmark.Get(context.Background(), name)
		if err != nil {
			result = multierror.Append(result, errors.Wrapf(err, "no such bookmark: %s", name))
			continue
		}
		if err := open.Run(b.URL); err != nil {
			o.ui.ErrPrintln("failed to open browser: %s", err)
			return 1
		}
	}

	if result != nil {
		o.ui.ErrPrintln(result)
		return 1
	}

	return 0
}

func (o *OpenCommand) Synopsis() string {
	return "open the url of a bookmark you selected"
}

func newOpenCommand() (cli.Command, error) {
	return &OpenCommand{ui: ui}, nil
}
