package main

import (
	"bytes"
	"testing"
)

func setupForCommandTest(t *testing.T) func() {
	old := repo
	repo = &Repository{
		Bookmark: &jsonBookmarkRepository{},
	}
	return func() {
		repo = old
	}
}

func TestAddCommand(t *testing.T) {
	dummyUI := &baseUI{
		Writer:    new(bytes.Buffer),
		ErrWriter: new(bytes.Buffer),
		Reader:    new(bytes.Buffer),
	}
	cmd := &AddCommand{
		ui: dummyUI,
	}

	cases := map[string]struct {
		in     []string
		hasErr bool
	}{
		"can add the URL of Google as named 'google'": {[]string{"https://google.com", "google"}, false},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			cleanup := setupForCommandTest(t)
			defer cleanup()

			code := cmd.Run(c.in)
			if c.hasErr {
				if code != 0 {
					t.Errorf("expected normal status code, but got abnormal code: %d", code)
				}
			} else {
				if code == 0 {
					t.Error("expected abnormal status code, but got normal code")
				}
			}
		})
	}
}
