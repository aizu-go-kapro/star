package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestListCommand(t *testing.T) {
	out := new(bytes.Buffer)
	dummyUI := &baseUI{
		writer:    out,
		errWriter: new(bytes.Buffer),
		reader:    new(bytes.Buffer),
	}
	cmd := &ListCommand{
		ui: dummyUI,
	}

	t.Run("display nothing when number of bookmarks is 0", func(t *testing.T) {
		_, cleanup := setupForCommandTest(t)
		defer cleanup()
		code := cmd.Run(nil)
		if code != 0 {
			t.Fatalf("expected normal status code, but got abnormal code: %d", code)
		}

		if out.Len() != 0 {
			t.Errorf("expected empty string, but got not empty string: %s", out.String())
		}
	})

	t.Run("display bookmarks which is stored in the DB", func(t *testing.T) {
		_, cleanup := setupForCommandTest(t)
		defer cleanup()

		// setup: load testdata

		b, err := ioutil.ReadFile(filepath.Join("testdata", "list.json"))
		if err != nil {
			t.Fatalf("failed to open testdata: %s", err)
		}

		var db DB
		if err := json.Unmarshal(b, &db); err != nil {
			t.Fatalf("failed to unmarshal testdata: %s", err)
		}
		repo = &Repository{
			Bookmark: newJSONBookmarkRepository(&db),
		}

		// do

		code := cmd.Run(nil)
		if code != 0 {
			t.Fatalf("expected normal status code, but got abnormal code: %d", code)
		}

		if !strings.Contains(out.String(), "google") {
			t.Errorf("expected output contains '%s', but missing", "google")
		}
		if !strings.Contains(out.String(), "https://google.com") {
			t.Errorf("expected output contains '%s', but missing", "https://google.com")
		}
	})
}
