package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var once sync.Once
var repo *Repository

func Init() {
	once.Do(func() {
		var err error
		repo, err = NewRepository()
		if err != nil {
			panic(err)
		}
	})
}

type DB struct {
	Bookmarks []*Bookmark `json:"bookmarks"`
}

type Bookmark struct {
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository struct {
	Bookmark BookmarkRepository
}

type BookmarkRepository interface {
	Add(context.Context, *Bookmark) error
	List(context.Context) ([]*Bookmark, error)
	Update(context.Context, *Bookmark) error
	Delete(context.Context, *Bookmark) error
}

type jsonBookmarkRepository struct {
	bookmarks []*Bookmark
}

func (j *jsonBookmarkRepository) Add(ctx context.Context, b *Bookmark) error {
	for _, e := range j.bookmarks {
		if e.Name == b.Name {
			return errors.New("Already exist bookmark name")
		}
	}
	j.bookmarks = append(j.bookmarks, b)
	return nil
}

func (j *jsonBookmarkRepository) List(_ context.Context) ([]*Bookmark, error) {
	return j.bookmarks, nil
}

func (j *jsonBookmarkRepository) Update(_ context.Context, b *Bookmark) error {
	for i, bookmark := range j.bookmarks {
		if bookmark.Name == b.Name {
			j.bookmarks[i] = b
			return nil
		}
	}
	return fmt.Errorf("Not found %s", b.Name)
}

func (j *jsonBookmarkRepository) Delete(_ context.Context, b *Bookmark) error {
	n, err := j.findBookmark(b)
	if err != nil {
		return err
	}
	switch {
	case n == 0 && len(j.bookmarks) == 1:
		j.bookmarks = []*Bookmark{}
	case n == 0:
		j.bookmarks = j.bookmarks[n+1:]
	case n == len(j.bookmarks)-1:
		j.bookmarks = j.bookmarks[:n]
	default:
		j.bookmarks = append(j.bookmarks[:n], j.bookmarks[n+1:]...)
	}
	return nil
}

func (j *jsonBookmarkRepository) findBookmark(b *Bookmark) (int, error) {
	for i := range j.bookmarks {
		if b.Name == j.bookmarks[i].Name {
			return i, nil
		}
	}
	return 0, errors.Errorf("no such named bookmark: %s", b.Name)
}

func NewRepository() (*Repository, error) {
	return newJSONRepository()
}

func newJSONRepository() (*Repository, error) {
	f, err := os.Open("in.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to load JSON database")
	}
	defer f.Close()

	var db DB
	if err := json.NewDecoder(f).Decode(&db); err != nil {
		return nil, errors.Wrap(err, "failed to decode JSON database")
	}

	return &Repository{
		Bookmark: &jsonBookmarkRepository{bookmarks: db.Bookmarks},
	}, nil
}
