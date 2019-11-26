#!/bin/sh
# This script consolidates the conversion of go comments to markdown
# performed by go2md.go into one single markdown file
# and rerenders the markdown file whenever
# changes are detected.

GIT_HUB_ROOT=https://github.com/egarbarino/go-by-assertion/tree/master/
DST=/mnt/d/GDrive/garba/draft/general/go-by-assertion/go-by-assertion.md
CONVERTER="go run src/main/go2md.go"
mkdir -p build
while true
do
  FILES=$(cat files | tr '\n' ' ')
  cat header.md > build/go-by-assertion.md 
  $CONVERTER -r $GIT_HUB_ROOT $FILES >> build/go-by-assertion.md 
  pandoc -s -S --toc build/go-by-assertion.md -o build/go-by-assertion.html
  cp build/go-by-assertion.md $DST 
  inotifywait -r src files header.md
done
