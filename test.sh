#!/bin/sh
FILES=$(find . -name "*_test.go")
for file in $FILES
do
  echo $file
  go test $file
done
        
        
