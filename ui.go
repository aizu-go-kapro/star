package main

import (
	"fmt"
	"io"
	"os"
)

var ui UI
var defaultUI = &baseUI{
	Writer:    os.Stdout,
	ErrWriter: os.Stderr,
	Reader:    os.Stdin,
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
}

type baseUI struct {
	Writer    io.Writer
	ErrWriter io.Writer
	Reader    io.Reader
}

func (u *baseUI) Println(a ...interface{}) {
	fmt.Fprintln(u.Writer, a...)
}

func (u *baseUI) ErrPrintln(a ...interface{}) {
	fmt.Fprintln(u.ErrWriter, a...)
}
