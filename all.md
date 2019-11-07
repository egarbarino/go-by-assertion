---
title: Go Examples By Unit Testing
author: Ernesto Garbarino
date: 2019-11-05
---

[Source](github.com/src/flowcontrol/flowcontrol_test.go)

``` go
var actions = ""
func addAction(action string) {
	actions = actions + action
}
func trackDefer() {
	addAction("1")
	defer addAction("[d1]")
	addAction("2")
	defer addAction("[d2]")
	addAction("3")
	defer addAction("[d3]")
	addAction("4")
	if true {
		return
	}
	defer addAction("[d4]")
	addAction("5")
}
trackDefer()
"1234[d3][d2][d1]" ⇔ actions
```

[Source](github.com/src/functions/functions_test.go)

``` go
var counter = 0
incrementer := func() int {
	counter++
	return counter
}
1 ⇔ incrementer()
counter++
3 ⇔ incrementer()
incrementer()
5 ⇔ incrementer()
```

[Source](github.com/src/basictypes/basictypes_test.go)

``` go
// Constants are immutable
const x byte = 5
const y = 7
// n++ // Won't compile
// Array-like like Slices are mutable and cannot be made constants
// const slice = []byte{'a', 'b', 'c'} // Won't compile
```

[Source](github.com/src/arrays/arrays_test.go)

``` go
array := [3]byte{'a', 'b', 'c'}
// Copying an array produces a brand new one
array2 := array
array2[0] = 'x'
byte('a') ⇔ array[0] // Rather than x
// Functions receive an array copy, rather than a reference
changeFirstElement(array)
byte('a') ⇔ array[0] // Rather than y
func changeFirstElement(array [3]byte) {
	array[0] = 'y'
}
```

[Source](github.com/src/slices/slices_test.go)

``` go
slice := []byte{'a', 'b', 'c'}
// A copy of a slice still points to the original slice
slice2 := slice
slice2[0] = 'x'
byte('x') ⇔ slice[0] // Rather than a
// Functions receive an slice pointer, rather than a copy
changeFirstElement(slice)
byte('y') ⇔ slice[0] // Rather than a or x
// A slice's slice still points to the original slice
lastTwoLetters := slice[1:]
lastTwoLetters[0] = 'z'
byte('z') ⇔ slice[1] // Rather than b
func changeFirstElement(slice []byte) {
	slice[0] = 'y'
}
```

[Source](github.com/src/strings/strings_test.go)

``` go
"Hello-world" ⇔ strings.Join([]string{"Hello", "world"}, "-")
"hihihi" ⇔ strings.Repeat("hi", 3)
```

[Source](github.com/src/strings/stringformatting_test.go)

``` go
"17" ⇔ fmt.Sprintf("%d", 17)             // Decimal
"11" ⇔ fmt.Sprintf("%x", 17)             // Hexadecimal
"21" ⇔ fmt.Sprintf("%o", 17)             // Octal
"10001" ⇔ fmt.Sprintf("%b", 17)          // Binary
"10001" ⇔ fmt.Sprintf("%b", 17)          // Binary
"1.750000" ⇔ fmt.Sprintf("%f", 1.75)     // Floating Point
"1.75" ⇔ fmt.Sprintf("%.2f", 1.75)       // Floating Point
"true" ⇔ fmt.Sprintf("%t", true)         // Boolean
"😊" ⇔ fmt.Sprintf("%c", '😊'))             // Rune (Unicode point
"hello" ⇔ fmt.Sprintf("%s", "hello"))     // Rune (Unicode point
"    o" ⇔ fmt.Sprintf("%5s", "o")        // Right Alignment
"h    " ⇔ fmt.Sprintf("%-5s", "h")       // Left Alignment
"'😊'" ⇔ fmt.Sprintf("%q", '😊')           // Quoted Rune
"'😊'" ⇔ fmt.Sprintf("%q", '😊')           // Quoted Rune
"\"hello\"" ⇔ fmt.Sprintf("%q", "hello") // Quoted String
"c64" ⇔ fmt.Sprintf("%v%v", "c", 64)     // Default String formatting
"int" ⇔ fmt.Sprintf("%T", 64)            // Inferred Value's Type
"%" ⇔ fmt.Sprintf("%%")                  // Literal Percent Sign
```

[Source](github.com/src/maps/maps_test.go)

``` go
// Set up Initial Map
var datesOfBirth = map[string]int{
	"Newton":   1643,
	"Faraday":  1791,
	"Einstein": 1879,
}
// Test key membership
_, ok := datesOfBirth["Newton"]
true ⇔ ok
_, ok2 := datesOfBirth["Curie"]
false ⇔ ok2
// Get all keys
var keys = ""
for key := range datesOfBirth {
	keys = keys + key
}
true ⇔ strings.Contains(keys, "Einstein")
true ⇔ strings.Contains(keys, "Newton")
true ⇔ strings.Contains(keys, "Faraday")
// Get all values
var values = ""
for _, value := range datesOfBirth {
	values = values + strconv.Itoa(value)
}
true ⇔ strings.Contains(values, "1879")
true ⇔ strings.Contains(values, "1643")
true ⇔ strings.Contains(values, "1791")
```

[Source](github.com/src/pointers/pointers_test.go)

``` go
var n int = 0
var n1, n2, n3 = &n, &n, &n
0 ⇔ n
0 ⇔ *n1
0 ⇔ *n2
0 ⇔ *n3
*n2++ // Any pointer to n (could have been n1 or n3)
1 ⇔ n
1 ⇔ *n1
1 ⇔ *n2
1 ⇔ *n3
```

[Source](github.com/src/structs/main/structs_test.go)

``` go
type Employee struct {
	Name    string
	Salary  int
	Address struct {
		Number     int
		StreetName string
		Postcode   string
	}
}
e1 := Employee{
	Name:   "Jazmin",
	Salary: 2000,
	Address: struct {
		Number     int
		StreetName string
		Postcode   string
	}{
		Number:     101,
		StreetName: "Buckhurst",
		Postcode:   "IG9IJA",
	},
}
"IG9IJA" ⇔ e1.Address.Postcode
```

[Source](github.com/src/structs/main/vendor/model/structs_export.go)

``` go
package model
type secretAgent struct { // First letter lower case: Unexported!
	name   string
	salary int
}
type PublicEmployee struct { // First letter upper case: Exported!
	Name   string // exported
	Role   string // exported
	salary int    // unexported (lower case!)
}
```

[Source](github.com/src/structs/main/structs_export_test.go)

``` go
package main
import (
	"testing"
	"model"
	"github.com/stretchr/testify/assert"
)
//s = model.secretAgent{"Luigi", 1000} // Unexported
//e = model.PublicEmployee{"Jazmin", "Doctor", 20000} Salary unexported!
e := model.PublicEmployee{Name: "Jazmin", Role: "Doctor"}
"Jazmin" ⇔ e.Name
"Doctor" ⇔ e.Role
// 2000 ⇔ e.Salary Salary ununexported!
```

[Source](github.com/src/structs/main/methods_test.go)

``` go
func (e *Employee) SucceedToIncreaseSalaryBy(extraSalary int) int {
	e.Salary = e.Salary + extraSalary // Passed value actually changed!
	return e.Salary
}
e := Employee{"Jazmin", 2000}
// Value-wise Method Call
var newSalary = e.FailToIncreaseSalaryBy(1000)
3000 ⇔ newSalary
2000 ⇔ e.Salary // Employee Salary not changed!
// Reference-wise Method Call (Side Effects!)
e.SucceedToIncreaseSalaryBy(1000)
3000 ⇔ e.Salary // Employee Salary changed!
```

[Source](github.com/src/structs/main/methods_inner_test.go)

``` go
func (s StringAlias) ToUpperCase() string {
	return strings.ToUpper(string(s))
}
greeting := StringAlias("hello")
"HELLO" ⇔ greeting.ToUpperCase()
```

[Source](github.com/src/files/files_test.go)

``` go
// First generate example file
contents := []byte("line1\nline2\nline3\n")
err := ioutil.WriteFile(tmpRoot+"lines.txt", contents, 0666)
if err != nil {
	t.Error("Can't open file: ", err)
}
// Open File For Reading
fileHandle, err := os.Open(tmpRoot + "lines.txt")
if err != nil {
	t.Error("Can't open file: ", err)
}
defer fileHandle.Close()
// Use Scanner
input := bufio.NewScanner(fileHandle)
var lineBuffer = ""
for input.Scan() {
	lineBuffer = lineBuffer + input.Text()
}
"line1line2line3" ⇔ lineBuffer
```

[Source](github.com/src/iostreams/iostreams.go)

``` go
func main() {
	// Read all Stdin
	data, _ := ioutil.ReadAll(os.Stdin)
	// Convert data to upper case and write to Stdout
	count, _ := os.Stdout.WriteString(strings.ToUpper(string(data)))
	// Write the number of bytes read to Stderr
	os.Stderr.WriteString(fmt.Sprintf("Bytes read: %d", count))
}
```

[Source](github.com/src/iostreams/iostreams_test.go)

``` go
// Specify command
cmd := exec.Command("go", "run", "iostreams.go")
// Specify stdin input
cmd.Stdin = strings.NewReader("hello")
// Declare a buffer to capture stdout from command
var stdout bytes.Buffer
cmd.Stdout = &stdout
// Declare a buffer to capture stderr from command
var stderr bytes.Buffer
cmd.Stderr = &stderr
// Run Command
err := cmd.Run()
if err != nil {
	t.Error("Error:", err)
}
// Verify stdout and stderr values
"HELLO" ⇔ stdout.String()
"Bytes read: 5" ⇔ stderr.String()
```
