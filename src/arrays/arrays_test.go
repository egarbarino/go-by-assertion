// # Arrays
//
// Arrays are immutable: once created, their size
// cannot be changed. They are passed by value by
// default.
//
// Ignore-On
package arrays

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Empty Array
// An empty arrays contains zero elements.
func Test_Init(t *testing.T) {

	var arrayEmpty1 [0]int
	arrayEmpty2 := [...]int{}

	// Assertions
	assert.Equal(t, arrayEmpty1, arrayEmpty2)
	assert.Equal(t, 0, len(arrayEmpty1))

}

// ## Blank Array
// A blank array is populated with the applicable
// zero element for the declared type.
func Test_Blank(t *testing.T) {

	var intArray [2]int
	var stringArray [2]string
	var fileArray [2]os.File

	// Assertions
	assert.Equal(t, 0, intArray[0])
	assert.Equal(t, 0, intArray[1])
	assert.Equal(t, "", stringArray[0])
	assert.Equal(t, "", stringArray[1])
	assert.Equal(t, os.File{}, fileArray[0])
	assert.Equal(t, os.File{}, fileArray[1])

}

// ## Array Literals
// Array literals are array declarations in which
// the space for the array is allocated and its
// elements are declared in one single statement.
func Test_Literals(t *testing.T) {

	// Explicit length
	arrayTwo := [2]string{"Einstein", "Newton"}
	// No need to specify length
	arrayThree := [...]string{"Einstein", "Newton", "Curie"}

	// Assertions
	assert.Equal(t, 2, len(arrayTwo))
	assert.Equal(t, 3, len(arrayThree))
}

// ## Multidimensional Arrays
// In multidimensional arrays, the outer elements are
// arrays (which in turn may declare further arrays!)

func Test_Multi(t *testing.T) {
	array := [2][2]string{
		{"A0", "B0"}, // Excel-like values
		{"A1", "B1"}, // for intuition
	}

	// Assertions
	assert.Equal(t, "A0", array[0][0])
	assert.Equal(t, "B0", array[0][1])
	assert.Equal(t, "A1", array[1][0])
	assert.Equal(t, "B1", array[1][1])
}

// ## Setting Elements' Values
// Elements are referred by their index, starting from zero.
func Test_SetValues(t *testing.T) {

	var array [2]string

	array[0] = "Einstein"
	array[1] = "Newton"

	// Assertions
	assert.Equal(t, "Einstein", array[0])
	assert.Equal(t, "Newton", array[1])
}

// ## Iteration (Classic)
// Arrays may be traversed using the traditional
// C-like approach: by getting their length and referencing
// each element by index.
func Test_Traverse(t *testing.T) {

	array := [3]byte{'a', 'b', 'c'}
	var result = ""
	var counter = 0
	for i := 0; i < len(array); i++ {
		counter++
		result = result + string(array[i])
	}

	// Assertions
	assert.Equal(t, 3, counter)
	assert.Equal(t, "abc", result)
}

// ## Iteration (Range)
//
// Please see [Control Flow](#control-flow)
//
// ## Pass By Value vs Pass By Reference
// Arrays are passed by value by default.
func Test_Reference(t *testing.T) {

	array := [3]byte{'a', 'b', 'c'}

	changeFirstElement := func(array [3]byte) {
		array[0] = 'x'
	}

	changeFirstElementByRef := func(array *[3]byte) {
		array[0] = 'x'
	}

	// Assertion #1
	// Copying an array produces a brand new one
	array2 := array
	array2[0] = 'x'
	assert.Equal(t, byte('a'), array[0]) // Rather than x

	// Assertion #2
	// Functions receive an array copy, rather than a reference
	changeFirstElement(array)
	assert.Equal(t, byte('a'), array[0]) // Rather than x

	// Assertion #3
	// Pass array's pointer
	changeFirstElementByRef(&array)
	assert.Equal(t, byte('x'), array[0]) // Rather than a

}
