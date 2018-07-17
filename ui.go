package main

import (
	"fmt"
	"io"
	"os"
)

var ui UI
var defaultUI = &baseUI{
	writer:    os.Stdout,
	errWriter: os.Stderr,
	reader:    os.Stdin,
}

// InitUI initializes the UI to output text.
// if passed UI is nil, defaultUI will be used.
func InitUI(u UI) {
	if u == nil {
		ui = defaultUI
		return
	}
	ui = u
}

type UI interface {
	Println(a ...interface{})
	ErrPrintln(a ...interface{})
	Writer() io.Writer
}

type baseUI struct {
	writer    io.Writer
	errWriter io.Writer
	reader    io.Reader
}

func (u *baseUI) Println(a ...interface{}) {
	fmt.Fprintln(u.writer, a...)
}

func (u *baseUI) ErrPrintln(a ...interface{}) {
	fmt.Fprintln(u.errWriter, a...)
}

func (u *baseUI) Writer() io.Writer {
	return u.writer
}
