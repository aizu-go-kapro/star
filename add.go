package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mitchellh/cli"
)

type AddCommand struct {
}

func (a *AddCommand) Help() string {
	return "Usage: star add <url> <name>"
}

func (a *AddCommand) Run(args []string) int {
	if len(args) != 2 {
		fmt.Println("Not match arguments")
		return 1
	}

	bookmark := &Bookmark{Name: args[0], URL: args[1], CreatedAt: time.Now()}
	if err := repo.Bookmark.Add(context.TODO(), bookmark); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func (a *AddCommand) Synopsis() string {
	return "add a bookmark"
}

func newAddCommand() (cli.Command, error) {
	return &AddCommand{}, nil
}
