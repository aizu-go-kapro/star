package main

import (
	"context"
	"testing"
	"time"
)

var _ BookmarkRepository = (*jsonBookmarkRepository)(nil)

func TestJSONBookmark(t *testing.T) {
	b := &Bookmark{Name: "bookmark", URL: "http://example.com", CreatedAt: time.Now()}
	repo := &jsonBookmarkRepository{}

	assertBookmarkLen := func(t *testing.T, expected int) {
		if len(repo.bookmarks) != expected {
			t.Errorf("expected number of %d bookmarks, but got %d", expected, len(repo.bookmarks))
		}
	}

	t.Run("Add adds a bookmark to the repository", func(t *testing.T) {
		assertBookmarkLen(t, 0)

		err := repo.Add(context.TODO(), b)
		if err != nil {
			t.Errorf("expected nil, but got error: %s", err)
		}

		assertBookmarkLen(t, 1)
	})

	// t.Run("List lists all bookmarks", func(t *testing.T) {
	// 	bookmarks, err := repo.List(context.Background())
	// 	if err != nil {
	// 		t.Fatalf("expected no errors, but got an error: %s", err)
	// 	}
	//
	// 	if len(bookmarks) != 1 {
	// 		t.Errorf("expected one bookmark, but got %d bookmarks", len(bookmarks))
	// 	}
	// })

	// t.Run("List lists all bookmarks", func(t *testing.T) {
	// 	bookmarks, err := repo.List(context.Background())
	// 	if err != nil {
	// 		t.Fatalf("expected no errors, but got an error: %s", err)
	// 	}
	//
	// 	if len(bookmarks) != 1 {
	// 		t.Errorf("expected one bookmark, but got %d bookmarks", len(bookmarks))
	// 	}
	// })
}
