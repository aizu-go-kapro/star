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

	cleanup := func(t *testing.T) {
		repo = &jsonBookmarkRepository{}
	}

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
		defer cleanup(t)

		addBookmarks(t, b)
		assertBookmarkLen(t, 1)
	})

	t.Run("Add cannot add values that have duplicated key", func(t *testing.T) {
		defer cleanup(t)

		addBookmarks(t, b)
		assertBookmarkLen(t, 1)
		err := repo.Add(context.Background(), b)
		if err == nil {
			t.Error("expected an error, but got nil")
		}
	})

	t.Run("List lists all bookmarks", func(t *testing.T) {
		defer cleanup(t)

		addBookmarks(t, b)
		assertBookmarkLen(t, 1)

		bookmarks, err := repo.List(context.Background())
		if err != nil {
			t.Fatalf("expected no errors, but got an error: %s", err)
		}

		if len(bookmarks) != 1 {
			t.Errorf("expected one bookmark, but got %d bookmarks", len(bookmarks))
		}
	})

	t.Run("Update fails if the value corresponding to passed key is not found", func(t *testing.T) {
		defer cleanup(t)

		err := repo.Update(context.Background(), b)
		if err == nil {
			t.Fatal("expected an error, but got no errors")
		}
	})

	t.Run("Update updates one bookmark from passed key (name)", func(t *testing.T) {
		defer cleanup(t)

		addBookmarks(t, b)
		assertBookmarkLen(t, 1)

		b2 := b
		b.CreatedAt = time.Now()
		b.URL = "https://updated-example.com"

		err := repo.Update(context.Background(), b2)
		if err != nil {
			t.Fatalf("expected no errors, but got an error: %s", err)
		}

		actual := repo.bookmarks[0]
		if *b2 != *actual {
			t.Errorf("expected that Update updates b by b2 (%v), but not equal (%v)", b2, actual)
		}
	})

	t.Run("Delete deletes one bookmark from passed key (name)", func(t *testing.T) {
		defer cleanup(t)

		addBookmarks(t, b)

		t.Run("delete one bookmark from the repository which has only one bookmark", func(t *testing.T) {
			err := repo.Delete(context.Background(), b)
			if err != nil {
				t.Fatalf("expected no errors, but got an error: %s", err)
			}

			assertBookmarkLen(t, 0)
		})

		t.Run("delete second bookmark from the repository which has tree bookmarks", func(t *testing.T) {
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
