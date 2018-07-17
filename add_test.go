package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"
)

func setupForCommandTest(t *testing.T) (io.Reader, func()) {
	old := repo
	repo = &Repository{
		Bookmark: &jsonBookmarkRepository{},
	}

	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err)
	}

	InitDB(f.Name())

	return f, func() {
		repo = old
		f.Close()
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
		"can add the URL of Google as named 'google'":    {[]string{"https://google.com", "google"}, false},
		"cannot add the URL of Google as named 'google'": {[]string{"google", "https://google.com"}, true},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			out, cleanup := setupForCommandTest(t)
			defer cleanup()

			code := cmd.Run(c.in)
			if c.hasErr {
				if code == 0 {
					t.Error("expected abnormal status code, but got normal code")
				}
				return
			} else {
				if code != 0 {
					t.Errorf("expected normal status code, but got abnormal code: %d", code)
				}
			}

			var db DB
			if err := json.NewDecoder(out).Decode(&db); err != nil {
				t.Fatalf("failed to decode test result: %s", err)
			}

			if len(db.Bookmarks) != 1 {
				t.Errorf("expected one bookmark is saved, but %d", len(db.Bookmarks))
			}
		})
	}
}
