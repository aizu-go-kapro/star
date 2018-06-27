package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/pkg/errors"
)

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
	// by tenntenn 重複管理はmap使ったほうがよいのでは？
	for _, e := range j.bookmarks {
		if e.Name == b.Name {
			return errors.New("Alrady exist bookmark name")
		}
	}
	j.bookmarks = append(j.bookmarks, b)
	return nil
}

func (j *jsonBookmarkRepository) List(context.Context) ([]*Bookmark, error) {
	panic("not implemented")
}

func (j *jsonBookmarkRepository) Update(context.Context, *Bookmark) error {
	panic("not implemented")
}

func (j *jsonBookmarkRepository) Delete(context.Context, *Bookmark) error {
	panic("not implemented")
}

func NewRepository() (*Repository, error) {
	return newJSONRepository()
}

// by tenntenn これはテスト用？ ファイル名が固定なのが気になる。
// テスト用ならばtestdata以下に移動したほうがいい。
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
