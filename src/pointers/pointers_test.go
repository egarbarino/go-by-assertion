// # Pointers
// Pointers are references to values, or more complex objects
// such as structs, maps, etc. Unlike C, there is no pointer
// arithmetic or the notion that a pointer to an array is equivalent
// to a pointer to its first element.
//
// Ignore-On
package pointers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Referencing and Dereferencing Pointers
// A pointer to a value is obtained by prefixing the value
// with `&`. If the value is a pointer, the value is obtained
// by prefixing it with `*`.
func Test_Pointer_Ref1(t *testing.T) {

	a := 0
	var b, c = &a, &a

	// Assertions #1
	assert.Equal(t, 0, a)    // a is the original value
	assert.NotEqual(t, 0, b) // b is a pointer to a. It is not zero
	assert.Equal(t, 0, *b)   // *b is the value of a
	assert.NotEqual(t, 0, c) // c is a pointer to a. It is not zero
	assert.Equal(t, 0, *c)   // *c is the value of a

	// Change
	*b++

	// Assertions #2
	assert.Equal(t, 1, a)  // a is the original value
	assert.Equal(t, 1, *b) // *b is the value of a
	assert.Equal(t, 1, *c) // *c is the value of a
}

// ## New: Obtaining a Pointer Rather Than a Value
// Variables allocated with the `new()` function are pointers rather than values.
func Test_Pointer_New(t *testing.T) {
	var n = new(int)

	// Assertions #1
	assert.Equal(t, 0, *n)
	assert.NotEqual(t, 0, n) // n is a pointer!

	// Change
	*n++

	// Assertions #2
	assert.Equal(t, 1, *n)
	assert.Equal(t, "0x", fmt.Sprintf("%v", n)[:2]) // n is a pointer!
}

// ## Pointer vs Value Behaviour
// The referencing and dereferencing operators assume that the target variable is
// a value, but not all objects behave like values.
func Test_Pointer_Pointer_vs_Value(t *testing.T) {
	// ------------------ Pointer Behaviour ------------------------

	// Slices
	slice := []int{0}
	sliceCopy := slice
	sliceCopy[0]++
	assert.Equal(t, 1, slice[0])
	assert.Equal(t, 1, sliceCopy[0])

	// Maps
	myMap := map[string]int{"a": 0}
	myMapCopy := myMap
	myMapCopy["a"]++
	assert.Equal(t, 1, myMap["a"])
	assert.Equal(t, 1, myMapCopy["a"])

	// Functions
	myFunc := func(n int) int { return n }
	myFuncCopy := myFunc
	assert.Equal(t, myFunc(1), myFuncCopy(1))

	// ------------------- Value Behaviour ------------------------

	// Structs
	type NumberWrapper struct {
		n int
	}
	myStruct := NumberWrapper{n: 0}
	myStructCopy := myStruct
	myStructCopy.n++
	assert.Equal(t, 0, myStruct.n)
	assert.Equal(t, 1, myStructCopy.n)

	// Arrays
	array := [1]int{0}
	arrayCopy := array
	arrayCopy[0]++
	assert.Equal(t, 0, array[0])
	assert.Equal(t, 1, arrayCopy[0])

	// Other built-in types
	// Strings, floats, integers, etc.
}
