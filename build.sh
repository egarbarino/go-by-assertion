#!/bin/sh
# This script consolidates the conversion of go comments to markdown
# performed by go2md.go into one single markdown file.

GIT_HUB_ROOT=https://github.com/egarbarino/go-examples/tree/master/
DST=/mnt/d/GDrive/garba/draft/general/go-by-assertion/go-by-assertion.md
CONVERTER="go run src/main/go2md.go"
FILES=$(cat files | tr '\n' ' ')
while true
do
  cat header.md > $DST 
  $CONVERTER -r $GIT_HUB_ROOT $FILES >> $DST 
  inotifywait -r src header.md
done
