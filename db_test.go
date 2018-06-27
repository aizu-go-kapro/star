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

	assertLen := func(t *testing.T, expected int) {
		bookmarks, err := repo.List(context.TODO())
		if err != nil {
			t.Fatalf("expected nil, but got error: %s", err)
		}

		if len(bookmarks) != 0 {
			t.Errorf("expected no bookmarks, but got %d", len(bookmarks))
		}
	}

	t.Run("Add adds a bookmark to the repository", func(t *testing.T) {
		assertLen(t, 0)

		err := repo.Add(context.TODO(), b)
		if err != nil {
			t.Errorf("expected nil, but got error: %s", err)
		}

		assertLen(t, 1)
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

	t.Run("List lists all bookmarks", func(t *testing.T) {
		bookmarks, err := repo.List(context.Background())
		if err != nil {
			t.Fatalf("expected no errors, but got an error: %s", err)
		}

		if len(bookmarks) != 1 {
			t.Errorf("expected one bookmark, but got %d bookmarks", len(bookmarks))
		}
	})
}
