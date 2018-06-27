package main

import "github.com/mitchellh/cli"

type UpdateCommand struct{}

// by tenntenn empty structなら func (Updatecommand) Help() string でも良さそう。
// 今後フィールドが増えるなら話は別。
func (u *UpdateCommand) Help() string {
	return "Usage: star update <url> <name>"
}

func (u *UpdateCommand) Run(args []string) int {
	panic("not implemented")
}

func (u *UpdateCommand) Synopsis() string {
	return "update a bookmark"
}

func newUpdateCommand() (cli.Command, error) {
	return &UpdateCommand{}, nil
}
