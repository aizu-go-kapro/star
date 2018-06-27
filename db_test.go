package main

import (
	"context"
	"testing"
	"time"
)

var _ BookmarkRepository = (*jsonBookmarkRepository)(nil)

func TestJSONBookmark_Add(t *testing.T) {
	repo := jsonBookmarkRepository{}
	bookmark := &Bookmark{"hoge", "https://hoge.example.com", time.Now()}

	if err := repo.Add(context.TODO(), bookmark); err != nil {
		t.Fatal(err)
	}
	if len(repo.bookmarks) != 1 {
		t.Fatalf("Not match array length")
	}
	if repo.bookmarks[0] != bookmark {
		t.Fatalf("Not Added bookmark")
	}

	if err := repo.Add(context.TODO(), bookmark); err == nil {
		t.Fatalf("Can't find already exist data")
	}
}
