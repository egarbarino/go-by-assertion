#!/bin/sh

DST=/mnt/d/GDrive/garba/article/general/go-examples/go-examples.md
CONVERTER="go run src/main/go2md.go"
FILES="src/flowcontrol/flowcontrol_test.go \
       src/functions/functions_test.go \
       src/basictypes/basictypes_test.go \
       src/arrays/arrays_test.go \
       src/slices/slices_test.go \
       src/strings/strings_test.go \
       src/strings/stringformatting_test.go \
       src/maps/maps_test.go \
       src/pointers/pointers_test.go \
       src/structs/main/structs_test.go \
       src/structs/main/vendor/model/structs_export.go \
       src/structs/main/structs_export_test.go \
       src/structs/main/methods_test.go \
       src/structs/main/methods_inner_test.go \
       src/files/files_test.go \
       src/iostreams/iostreams.go \
       src/iostreams/iostreams_test.go" 

  
echo "---" > all.md
echo "title: Go Examples By Unit Testing" >> all.md
echo "author: Ernesto Garbarino" >> all.md
echo "date: 2019-11-05" >> all.md
echo "---" >> all.md
$CONVERTER -r github.com/ $FILES >> all.md
cp all.md $DST
# less all.md
