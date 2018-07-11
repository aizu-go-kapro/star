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

	addBookmarks := func(t *testing.T, b ...*Bookmark) {
		for _, bookmark := range b {
			err := repo.Add(context.TODO(), bookmark)
			if err != nil {
				t.Errorf("expected nil, but got error: %s", err)
			}
		}
	}
	assertBookmarkContext := func(t *testing.T, expected *Bookmark) {
		if repo.bookmarks[len(repo.bookmarks)-1] != expected {
			t.Errorf("expected %s, but got %s", expected.Name, repo.bookmarks[len(repo.bookmarks)-1].Name)
		}
	}
	updateBookmark := func(t *testing.T, b ...*Bookmark) {
		for _, bookmark := range b {
			err := repo.Update(context.TODO(), bookmark)
			if err != nil {
				t.Errorf("expected nil, but got error: %s", err)
			}
		}
	}

	t.Run("Add adds a bookmark to the repository", func(t *testing.T) {
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
	t.Run("Update updates one bookmark from passed key (name)", func(t *testing.T) {
		err := repo.Update(context.Background(), b)
		if err != nil {
			t.Fatalf("expected no errors, but got an error: %s", err)
		}

		assertBookmarkLen(t, 0)
	})
	t.Run("ブックマークの要素数が変わっていないか確認", func(t *testing.T) {
		b2 := &Bookmark{Name: "bookmark2", URL: "http://foo.com", CreatedAt: time.Now()}
		assertBookmarkLen(t, 0)

		addBookmarks(t, b)

		updateBookmark(t, b2)

		assertBookmarkLen(t, 1)
	})
	t.Run("ブックマークの内容が変わっているか確認", func(t *testing.T) {
		b2 := &Bookmark{Name: "bookmark2", URL: "http://foo.com", CreatedAt: time.Now()}
		b3 := &Bookmark{Name: "bookmark3", URL: "http://bar.com", CreatedAt: time.Now()}
		b4 := &Bookmark{Name: "bookmark4", URL: "http://hoge.com", CreatedAt: time.Now()}
		assertBookmarkLen(t, 0)

		addBookmarks(t, b, b2, b3)

		Rebookmarks := repo.bookmarks

		updateBookmark(t, b4)

		var count = 2

		for i, b := range repo.bookmarks {
			if b.Name == Rebookmarks[i].Name {
				count--
			}
		}
		if count != 0 {
			t.Fatalf("expected the others are same, but exist not same")
		}
		for i, b := range repo.bookmarks {
			if b.Name == b4.Name {
				break
			}
			if i == len(repo.bookmarks)-1 {
				t.Fatalf("expected complete changing, but couldn't change: %s", b4.Name)
			}
		}
	})
	t.Run("Delete deletes one bookmark from passed key (name)", func(t *testing.T) {
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
