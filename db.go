package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	bookmarks sync.Map
}

func newJSONBookmarkRepository(db *DB) *jsonBookmarkRepository {
	var m sync.Map
	for _, b := range db.Bookmarks {
		_, ok := m.LoadOrStore(b.Name, b)
		if ok {
			panic(fmt.Sprintf("duplicated key found in the JSON DB: %s", b.Name))
		}
	}
	return &jsonBookmarkRepository{
		bookmarks: m,
	}
}

func (j *jsonBookmarkRepository) Add(ctx context.Context, b *Bookmark) error {
	_, ok := j.bookmarks.LoadOrStore(b.Name, b)
	if ok {
		return errors.New("Already exist bookmark name")
	}
	return nil
}

func (j *jsonBookmarkRepository) List(_ context.Context) ([]*Bookmark, error) {
	return j.slice(), nil
}

func (j *jsonBookmarkRepository) Update(_ context.Context, b *Bookmark) error {
	_, ok := j.bookmarks.LoadOrStore(b.Name, b)
	if !ok {
		return errors.Wrap(errNotFoundBookmark, "failed to find the bookmark specified by passed key")
	}
	return nil
}

func (j *jsonBookmarkRepository) Delete(_ context.Context, b *Bookmark) error {
	_, ok := j.bookmarks.Load(b.Name)
	if !ok {
		return errors.Wrap(errNotFoundBookmark, "failed to find the bookmark specified by passed key")
	}
	j.bookmarks.Delete(b.Name)
	return nil
}

func (j *jsonBookmarkRepository) slice() []*Bookmark {
	var b []*Bookmark
	j.bookmarks.Range(func(k, v interface{}) bool {
		b = append(b, v.(*Bookmark))
		return true
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i].Name < b[j].Name
	})
	return b
}

var errNotFoundBookmark = errors.New("no such named bookmark")

func notFoundBookmark(err error) bool {
	return errors.Cause(err) == errNotFoundBookmark
}

func NewRepository() (*Repository, error) {
	return newJSONRepository()
}

// TODO (@ktr0731)
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
		Bookmark: newJSONBookmarkRepository(&db),
	}, nil
}
