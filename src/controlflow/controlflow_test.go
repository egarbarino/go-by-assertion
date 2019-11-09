// # Control Flow
// Ignore-On
package flowcontrol

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## If-Then
// If statements don't use parenthesis in Go, unlike other C-like languages.
func Test_If(t *testing.T) {
	var counter = 0

	// Simple If
	if counter == 0 {
		counter++
	}

	// Assertions
	assert.Equal(t, 1, counter)
}

// ## If-Then-Else
func Test_IfThenElse(t *testing.T) {
	var counter = 0

	// Simple If
	if counter == 0 {
		counter++
	} else {
		counter--
	}

	// Assertions
	assert.Equal(t, 1, counter)
}

// ## If-Then-Else If
func Test_IfThenElseIf(t *testing.T) {
	var counter = 0

	// Simple If
	if counter == 0 {
		counter++
	} else if counter == 10 {
		counter = counter * 10
	} else if counter == 20 {
		counter = counter * 20
	} else {
		counter--
	}

	// Assertions
	assert.Equal(t, 1, counter)
}

// ## If-Scoped Variables
// An If statement may introduce variables to its scope.
func Test_If_Scope(t *testing.T) {

	var counter = 0

	// Assign
	if i := 5; counter == 0 {
		counter = counter + i // i in scope here
	} else {
		counter = counter - i // i in scope here
	}

	// --- i is not longer in scope here ---

	// Assertions
	assert.Equal(t, 5, counter)
}

// ## Switch (Normal)
func Test_Switch_Normal(t *testing.T) {

	var result = ""

	switch 1 + 1 {
	case 0:
		result = "zero"
	case 1:
		result = "one"
	case 2:
		result = "two"
	default: // Optional
		result = "some_number"
	}

	// Assertions
	assert.Equal(t, "two", result)
}

// ## Switch (Multi-Value Case)
func Test_Switch_Multi_Value(t *testing.T) {

	var result = ""

	// Just cases
	switch 2 * 3 {
	case 0, 1, 2, 3, 4, 5, 6:
		result = "between_zero_and_six"
	case 7:
		result = "seven"
	default: // Optional
		result = "some_number"
	}

	// Assertions
	assert.Equal(t, "between_zero_and_six", result)
}

// ## Switch (Expression Cases)
func Test_Switch_Expression_Cases(t *testing.T) {

	var result = ""

	switch { // No expression here
	case 2*3 == 1:
		result = "one"
	case 2*3 == 6:
		result = "six"
	default: // Optional
		result = "some_number"
	}

	// Assertions
	assert.Equal(t, "six", result)
}

// ## Switch (Type Cases)
func Test_Switch_Type_Cases(t *testing.T) {
	testType := func(i interface{}) string {
		switch t := i.(type) { // the t := assignment is optional
		case int:
			return "int"
		case string:
			return "string"
		default: // Optional
			return "other type: " + fmt.Sprintf("%T", t)
		}
	}

	// Assertions
	assert.Equal(t, "string", testType("hello"))
	assert.Equal(t, "int", testType(45))
	assert.Equal(t, "other type: float64", testType(1.53))
}

// ## Inifinite Loop (While True Loop)
func Test_While_True_Loop(t *testing.T) {

	var counter = 0

	for {
		counter = counter + 1
		if counter == 3 {
			break
		}
	}

	// Assertions
	assert.Equal(t, 3, counter)
}

// ## While Loop
func Test_While_Loop(t *testing.T) {

	var counter = 0

	for counter < 3 {
		counter = counter + 1
	}

	// Assertions
	assert.Equal(t, 3, counter)
}

// ## Iterate Over Array Elements
func Test_Range_Array_Loop(t *testing.T) {

	var indexSum = 0
	var sum = 0

	for index, currentValue := range []int{1, 2, 3} {
		indexSum = indexSum + index // 0 + 1 + 2 = 3
		sum = sum + currentValue    // 1 + 2 + 3 = 6
	}

	// Assertions
	assert.Equal(t, 3, indexSum)
	assert.Equal(t, 6, sum)
}

// ## Iterate Over Array Elements But Use Only Index
func Test_Range_Array_Loop_Ignore_Index(t *testing.T) {

	var indexSum = 0

	for index := range []int{50, 999, 300} {
		indexSum = indexSum + index // 0 + 1 + 2 = 3
	}

	// Assertions
	assert.Equal(t, 3, indexSum)
}

// ## Iterate Over Array Elements But Ignore Index
func Test_Range_Array_Loop_Ignore_Index_Blank_Identifier(t *testing.T) {

	var sum = 0

	for _, currentValue := range []int{1, 2, 3} {
		sum = sum + currentValue // 1 + 2 + 3 = 6
	}

	// Assertions
	assert.Equal(t, 6, sum)
}

// ## Iterate Over Keys and Values of a Map
func Test_Range_Map_Loop(t *testing.T) {

	var keys, values string

	for k, v := range map[string]string{"A": "(Argentina)", "B": "(Brazil)"} {
		keys = keys + k
		values = values + v
	}

	// Assertions
	assert.Equal(t, "AB", keys)
	assert.Equal(t, "(Argentina)(Brazil)", values)
}

// ## Iterate Over the Unicode Characters of a String
func Test_Range_String_Loop(t *testing.T) {
	var word string

	for _, v := range "😊 olleh" {
		word = string(v) + word
	}

	// Assertions
	assert.Equal(t, "hello 😊", word)
}

// ## Iterate Over The Bytes of a String
func Test_For_String_Bytes_Loop(t *testing.T) {

	var reversedHello = "😊 olleh"
	var word string

	for i := 0; i < len(reversedHello); i++ {
		word = string(reversedHello[i]) + word
	}

	// Assertions
	assert.Equal(t, "hello \u008a\u0098\u009fð", word)
}

// ## Regular Foor Loop
func Test_Regular_For_Loop(t *testing.T) {

	var counter = 0

	for i := 0; i < 3; i++ {
		counter = counter + i // 0 + 1 + 2 = 3
	}

	// Assertions
	assert.Equal(t, 3, counter)
}

// ## Defer

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
func Test_Defer(t *testing.T) {
	trackDefer()

	// Assertions
	assert.Equal(t, "1234[d3][d2][d1]", actions)
}
