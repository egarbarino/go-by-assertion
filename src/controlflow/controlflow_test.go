// # Control Flow
// Ignore-On
package flowcontrol

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## If-Then
// The syntax is `if EXPR {...}`.
// _If statements_ don't use parentheses in Go, unlike other C-like languages.
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
// The syntax is `if EXPR {...} else {...}`.
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
// The syntax is `if EXPR {...} else if EXPR {...}`. A final `else {...}`, as
// in the example below, may also be included.
func Test_IfThenElseIf(t *testing.T) {
	var counter = 0

	// Simple If
	if counter == 0 {
		counter++
	} else if counter == 10 {
		counter = counter * 10
	} else if counter == 20 {
		counter = counter * 20
	} else { // Optional
		counter--
	}

	// Assertions
	assert.Equal(t, 1, counter)
}

// ## If-Scoped Variables
// An _if statement_, in the form `if VAR_DECLARATION ; EXPR {...}`,
// introduces variables to its scope.
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

// ## Switch
// The regular `switch EXPR {...}` statement evaluates an expression `EXPR` and
// executes the _first_ case
// whose value matches the reduction of said expression. Unlike other languages,
// the cases do not "cascade" and there is no need to terminate them with a `break`
// statement or similar.

// The `default` keyword is used to label the case to be selected when all the
// other ones prefixed with `case` fail to match.
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
// The multi-value switch allows to match multiple values within
// a single case statement.
func Test_Switch_Multi_Value(t *testing.T) {
	var result = ""

	// Just cases
	switch 2 * 3 {
	case 0, 1, 2, 3, 4, 5, 6: // multiple values!
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
// This is a type of switch statement in which the cases contain boolean expressions
// rather then values. The first case whose expression reduces to true
// will be executed.
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

// A more crude example is presented below:
func Test_Switch_Expression_Cases_2(t *testing.T) {
	var result = ""

	switch { // No expression here
	case false:
		result = "false1"
	case false:
		result = "false2"
	case true:
		result = "true1"
	default: // Optional
		result = "true2"
	}

	// Assertions
	assert.Equal(t, "true1", result)
}

// ## Switch (Type Cases)
// This type of switch statement facilitates implementing code based on the different
// possible types of a given value.
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

// ## Infinite Loop (While True Loop)
// An infinite loop, in the form `for {...}` never terminates
// and must be short-circuited explicitly; for
// example, using `break` as in the example below.
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
// A while loop has the form `for EXPR {...}` where the body is executed
// repeatedly as long as `EXPR` is true.
func Test_While_Loop(t *testing.T) {

	var counter = 0

	for counter < 3 {
		counter = counter + 1
	}

	// Assertions
	assert.Equal(t, 3, counter)
}

// ## Iteration Over Array Elements
// This is accomplished by using the `for index, value := range ARRAY {...}` syntax
// where index is the array index starting from zero.
func Test_Range_Array_Loop(t *testing.T) {

	var indexSum = 0
	var sum = 0

	for index, currentValue := range [3]int{1, 2, 3} {
		indexSum = indexSum + index // 0 + 1 + 2 = 3
		sum = sum + currentValue    // 1 + 2 + 3 = 6
	}

	// Assertions
	assert.Equal(t, 3, indexSum)
	assert.Equal(t, 6, sum)
}

// ## Iteration Over Array Elements Using Only Index
// This is like the regular iteration over an array except that the second
// value is ignored in the assignment: `index := range ARRAY {...}`
func Test_Range_Array_Loop_Ignore_Index(t *testing.T) {

	var sum = 0
	var numbers = [3]int{1, 2, 3}

	for index := range numbers {
		sum = sum + numbers[index] // array elements accessed by index!
	}

	// Assertions
	assert.Equal(t, 6, sum)
}

// ## Iteration Over Array Elements Ignoring Index
// In this case, the index value is explicitly ignored by using the underscore `_`
// _blank identifier_ symbol.
func Test_Range_Array_Loop_Ignore_Index_Blank_Identifier(t *testing.T) {
	var sum = 0

	for _, currentValue := range [3]int{1, 2, 3} {
		sum = sum + currentValue // 1 + 2 + 3 = 6
	}

	// Assertions
	assert.Equal(t, 6, sum)
}

// ## Iteration Over Keys and Values of a Map
// This is achieved using the `key, value := range MAP {...}` syntax.
func Test_Range_Map_Loop(t *testing.T) {
	var keys, values string

	for k, v := range map[string]string{"A": "Argentina", "B": "Brazil"} {
		keys = keys + k
		values = values + v
	}

	// Assertions
	assert.Equal(t, true, strings.Contains(keys, "A"))
	assert.Equal(t, true, strings.Contains(keys, "B"))
	assert.Equal(t, true, strings.Contains(values, "Argentina"))
	assert.Equal(t, true, strings.Contains(values, "Brazil"))
}

// ## Iteration Over the Unicode Characters of a String
// By prefixing an string with `range`, the `for` loop will iterate through the
// string's Unicode characters as opposed to its raw bytes.
func Test_Range_String_Loop(t *testing.T) {
	var word string

	for _, v := range "😊 olleh" {
		word = string(v) + word
	}

	// Assertions
	assert.Equal(t, "hello 😊", word)
}

// ## Iteration Over The Bytes of a String
// This is like iterating over an array in a regular C-like language.
func Test_For_String_Bytes_Loop(t *testing.T) {

	var reversedHello = "😊 olleh"
	var word string

	for i := 0; i < len(reversedHello); i++ {
		word = string(reversedHello[i]) + word
	}

	// Assertions
	assert.Equal(t, "hello \u008a\u0098\u009fð", word)
}

// ## Regular For Loop
// The regular for loop is exactly the same as that found in C-like
// languages except that no parentheses are used and that the
// variable(s) are initialised with `:=` rather than `=`.
func Test_Regular_For_Loop(t *testing.T) {

	var counter = 0

	for i := 0; i < 3; i++ {
		counter = counter + i // 0 + 1 + 2 = 3
	}

	// Assertions
	assert.Equal(t, 3, counter)
}

// ## Defer
// The `defer STATEMENT` syntax is used to enforce the execution of the
// statement `STATEMENT`
// before the function under which the declaration appears is exited.
// The execution order is that of a stack: last in, first out (LIFO).
func Test_Defer(t *testing.T) {
	var actions = ""

	addAction := func(action string) {
		actions = actions + action
	}

	trackDefer := func() {
		addAction("1")
		defer addAction("[d1]")
		addAction("2")
		defer addAction("[d2]")
		addAction("3")
		defer addAction("[d3]")
		addAction("4")
	}

	actions = actions + "START-"
	trackDefer()
	actions = actions + "-END"

	// Assertions
	assert.Equal(t, "START-1234[d3][d2][d1]-END", actions)
}
