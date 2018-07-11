package main

import (
	"context"
	"io/ioutil"
	"testing"
	"time"
)

var _ BookmarkRepository = (*jsonBookmarkRepository)(nil)

func setupForDBTest(t *testing.T) func() {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err)
	}
	dbPath = f.Name()
	return func() {
		f.Close()
	}
}

func TestJSONBookmark(t *testing.T) {
	b := &Bookmark{Name: "bookmark", URL: "http://example.com", CreatedAt: time.Now()}
	repo := &jsonBookmarkRepository{}

	assertBookmarkLen := func(t *testing.T, expected int) {
		if len(repo.bookmarks) != expected {
			t.Errorf("expected number of %d bookmarks, but got %d", expected, len(repo.bookmarks))
		}
	}

	addBookmarks := func(t *testing.T, b ...*Bookmark) {
		for _, bookmark := range b {
			err := repo.Add(context.TODO(), bookmark)
			if err != nil {
				t.Errorf("expected nil, but got error: %s", err)
			}
		}
	}

	t.Run("Add adds a bookmark to the repository", func(t *testing.T) {
		cleanup := setupForDBTest(t)
		defer cleanup()

		assertBookmarkLen(t, 0)

		addBookmarks(t, b)

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

	t.Run("Delete deletes one bookmark from passed key (name)", func(t *testing.T) {
		t.Run("delete one bookmark from the repository which has only one bookmark", func(t *testing.T) {
			cleanup := setupForDBTest(t)
			defer cleanup()

			err := repo.Delete(context.Background(), b)
			if err != nil {
				t.Fatalf("expected no errors, but got an error: %s", err)
			}

			assertBookmarkLen(t, 0)
		})

		t.Run("delete second bookmark from the repository which has tree bookmarks", func(t *testing.T) {
			cleanup := setupForDBTest(t)
			defer cleanup()

			b2 := &Bookmark{Name: "bookmark2", URL: "http://foo.com", CreatedAt: time.Now()}
			b3 := &Bookmark{Name: "bookmark3", URL: "http://bar.com", CreatedAt: time.Now()}
			addBookmarks(t, b, b2, b3)

			assertBookmarkLen(t, 3)

			err := repo.Delete(context.Background(), b2)
			if err != nil {
				t.Fatalf("expected no errors, but got an error: %s", err)
			}

			assertBookmarkLen(t, 2)
			if repo.bookmarks[0] != b {
				t.Errorf("expected %v, but got %v", b, repo.bookmarks[0])
			}
			if repo.bookmarks[1] != b3 {
				t.Errorf("expected %v, but got %v", b3, repo.bookmarks[1])
			}
		})
	})
}
