package main

import (
	"testing"
)

var _ BookmarkRepository = (*jsonBookmarkRepository)(nil)

func TestJSONBookmark_Add(t *testing.T) {}
