# star
Team 音ゲーマー project

ブックマークをコマンドライン上で管理するためのツール

## usage

``` bash
# add three bookmarks
$ star add https://github.com github
$ star add http://syfm.hatenablog.com blog
$ star add https://godoc.org/fmt fmt

# list ups
$ star list
fmt    https://godoc.org/fmt
github https://github.com
blog   http://syfm.hatenablog.com

# delete a bookmark
$ star delete github
$ star list
fmt  https://godoc.org/fmt
blog http://syfm.hatenablog.com

# update a bookmark
$ star update blog https://syfm.hatenablog.com
$ star list
fmt    https://godoc.org/fmt
hatebu https://b.hatena.ne.jp

# open a bookmark
$ star open fmt

# (example) open all bookmarks
$ star list | awk '{ print $1 }' | xargs star open

# open bookmarks with fuzzy finder
$ star open -f
```
