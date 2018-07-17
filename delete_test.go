package main

import (
	"bytes"
	"context"
	"testing"
	"time"
)

func TestDeleteCommand(t *testing.T) {
	dummyUI := &baseUI{
		writer:    new(bytes.Buffer),
		errWriter: new(bytes.Buffer),
		reader:    new(bytes.Buffer),
	}
	cmd := &DeleteCommand{
		ui: dummyUI,
	}

	t.Run("cannot delete when number of bookmarks is 0", func(t *testing.T) {
		_, cleanup := setupForCommandTest(t)
		defer cleanup()
		code := cmd.Run([]string{"google"})
		if code == 0 {
			t.Fatal("expected abnormal status code, but got normal code")
		}
	})

	// TODO: use interface and bytes.Buffer
	t.Run("can delete a bookmark", func(t *testing.T) {
		_, cleanup := setupForCommandTest(t)
		defer cleanup()
		err := repo.Bookmark.Add(context.Background(), &Bookmark{Name: "google", URL: "https://google.com", CreatedAt: time.Now()})
		if err != nil {
			t.Fatalf("failed to add a bookmark: %s", err)
		}
	})
}
