package main

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/mitchellh/cli"
)

type ListCommand struct {
	ui UI
}

func (l *ListCommand) Help() string {
	return "Usage: star list"
}

func (l *ListCommand) Run(args []string) int {
	bookmarks, err := repo.Bookmark.List(context.Background())
	if err != nil {
		l.ui.ErrPrintln("cannot list bookmarks: %s", err)
		return 1
	}

	w := tabwriter.NewWriter(l.ui.Writer(), 0, 0, 1, ' ', tabwriter.TabIndent)
	for _, b := range bookmarks {
		fmt.Fprintf(w, "%s\t%s\n", b.Name, b.URL)
	}
	if err := w.Flush(); err != nil {
		l.ui.ErrPrintln("cannot display bookmarks: %s", err)
		return 1
	}
	return 0
}

func (l *ListCommand) Synopsis() string {
	return "list ups all bookmarks"
}

func newListCommand() (cli.Command, error) {
	return &ListCommand{ui: ui}, nil
}
