package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/mattn/go-pipeline"
	"github.com/mitchellh/cli"
	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
)

type OpenCommand struct {
	ui UI

	f              *flag.FlagSet
	useFuzzySearch bool
}

func (o *OpenCommand) Help() string {
	out := new(bytes.Buffer)
	o.f.SetOutput(out)
	o.f.Usage()
	return fmt.Sprintf("Usage: star open <name>\n%s", out.String())
}

func (o *OpenCommand) Run(args []string) int {
	o.f.Parse(args)
	args = o.f.Args()

	if o.useFuzzySearch {
		var err error
		args, err = o.runFuzzyOpen(o.f.Args())
		if err != nil {
			o.ui.ErrPrintln(err)
			return 1
		}
	}

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

// TODO: other finders
var fzfPathEnvName = "FZF_PATH"

func (o *OpenCommand) runFuzzyOpen(args []string) ([]string, error) {
	var cmdPath string
	switch {
	case os.Getenv(fzfPathEnvName) != "":
		cmdPath = os.Getenv(fzfPathEnvName)
	case isAizuEnv():
		// TODO
		cmdPath = "/home/student/s1230022/.fzf/bin/fzf"
	default:
		var err error
		cmdPath, err = exec.LookPath("fzf")
		if err != nil {
			return nil, errors.New("you need to install fzf to executable paths, or set $FZF_PATH env to use fuzzy search mode")
		}
	}

	out, errOut := new(bytes.Buffer), new(bytes.Buffer)
	cmd := &ListCommand{ui: &baseUI{writer: out, errWriter: errOut}}
	if code := cmd.Run(nil); code != 0 {
		return nil, errors.Errorf("failed to get all bookmarks: %s", errOut.String())
	}

	b, err := pipeline.Output(
		[]string{"echo", out.String()},
		[]string{cmdPath, "-m"},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute fzf")
	}

	fmt.Println(string(b))

	sp := strings.Split(strings.TrimSpace(string(b)), "\n")
	bookmarks := make([]string, 0, len(sp))
	for _, s := range sp {
		bookmarks = append(bookmarks, strings.SplitN(s, " ", 2)[0])
	}

	return bookmarks, nil
}

func newOpenCommand() (cli.Command, error) {
	cmd := &OpenCommand{ui: ui}
	cmd.f = flag.NewFlagSet("open", flag.ExitOnError)
	cmd.f.BoolVar(&cmd.useFuzzySearch, "f", false, "use fuzzy search")
	cmd.f.Usage = cmd.f.PrintDefaults
	return cmd, nil
}

func isAizuEnv() bool {
	if runtime.GOOS != "solaris" {
		return false
	}

	cmd := exec.Command("defaultdomain")
	out := new(bytes.Buffer)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return false
	}

	if strings.Contains(out.String(), "u-aizu.ac.jp") {
		return true
	}
	return false
}
