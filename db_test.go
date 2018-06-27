package main

import (
	"context"
	"testing"
	"time"
)

var _ BookmarkRepository = (*jsonBookmarkRepository)(nil)

func TestJSONBookmark_Add(t *testing.T) {}
func TestJSONBookmark_List(t *testing.T) {
	b := &Bookmark{Name: "bookmark", URL: "http://example.com", CreatedAt: time.Now()}
	repo := &jsonBookmarkRepository{}

	assertLen := func(t *testing.T, expected int) {
		bookmarks, err := repo.List(context.TODO())
		if err != nil {
			t.Fatalf("expected nil, but got error: %s", err)
		}

		if len(bookmarks) != 0 {
			t.Errorf("expected no bookmarks, but got %d", len(bookmarks))
		}
	}

	assertLen(t, 0)

	err := repo.Add(context.TODO(), b)
	if err != nil {
		t.Fatalf("expected nil, but got error: %s", err)
	}

	assertLen(t, 1)
}
