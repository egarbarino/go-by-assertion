// # Slices
//
// Slices are in principle "views" or "projections"
// over arrays created using the slice syntax, for example,
// `[3]string{"red","green","blue"}[0:2]` produces a _slice_
// that points to the `{"red","green"}` elements of the array.
// However, the idiomatic treatment of slices in Go is to always
// work with slices from the get go even when the initial slice
// could have been implemented as an array.
//
// The key message is that a slice is a type in itself
// rather than an array with a lesser number of elements.
//
// Ignore-On
package slices

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Slices vs Arrays
// This example shows that when we get an array's slice,
// a brand new type (hmm, a _slice_), is produced, rather
// than a smaller array.
func Test_Versus(t *testing.T) {
	var array = [3]string{"red", "green", "blue"}

	// Assertions
	assert.Equal(t, []string{"red", "green"}, array[0:2])     // Slice is a length-free type
	assert.NotEqual(t, [2]string{"red", "green"}, array[0:2]) // Slice is not a smaller array
}

// ## Empty Slice
// Unlike plain arrays, slices may be declared as empty
// but then populated using the `append()` function.
// The implementation is not, however, optimised for
// this use case: new arrays may be allocated
// depending on the slice's initial capacity.
func Test_Init(t *testing.T) {
	var sliceEmpty1 []int
	var sliceEmpty2 = make([]int, 0)
	var sliceEmpty3 = []int{}

	// Assertions #1
	// Prove slices have no elements
	assert.Equal(t, 0, len(sliceEmpty1))
	assert.Equal(t, 0, len(sliceEmpty2))
	assert.Equal(t, 0, len(sliceEmpty3))

	// Assertions #2
	// Append one element
	assert.Equal(t, []int{5}, append(sliceEmpty1, 5))
	assert.Equal(t, []int{5}, append(sliceEmpty2, 5))
	assert.Equal(t, []int{5}, append(sliceEmpty3, 5))
}

// ## Make (Slice Allocation)
// The `make()` function helps allocate a new slice by
// specifying the initial slice's length and capacity.
// The specific syntax for slices is
// `make([]<TYPE>, <LENGTH>, [<CAPACITY>])`. The functions
// `len()` and `cap()` are used to obtain a slice's length
// and capacity, respectively.
func Test_Make(t *testing.T) {
	var slice1 = make([]int, 2)    // Capacity omitted
	var slice2 = make([]int, 2, 2) // Equal length and capacity
	var slice3 = make([]int, 1, 2) // Capacity larger then length
	// var slice4 = make([]int, 3, 2) // Length larger than capacity (Compiler error)

	// Assertions #1
	assert.Equal(t, []int{0, 0}, slice1)
	assert.Equal(t, 2, len(slice1))
	assert.Equal(t, 2, cap(slice1))

	// Assertions #2
	assert.Equal(t, []int{0, 0}, slice2)
	assert.Equal(t, 2, len(slice2))
	assert.Equal(t, 2, cap(slice2))

	// Assertions #3
	assert.Equal(t, []int{0}, slice3)
	assert.Equal(t, 1, len(slice3))
	assert.Equal(t, 2, cap(slice3))

}

// ## Slice Literals
// A slice literal combines the allocation of a slice,
// and the setting of its elements, in one statement.
func Test_Literals(t *testing.T) {

	slice := []string{"Einstein", "Newton"}

	// Assertions
	assert.Equal(t, 2, len(slice))

}

// ## Setting Elements' Values
// Like in the case of arrays, elements are referenced
// by index, starting from zero.
func Test_Set_Values(t *testing.T) {

	slice := []byte{'a', 'b', 'c'}

	// Assertions
	assert.Equal(t, byte('a'), slice[0])
	slice[0] = 'x' // Change first element!
	assert.Equal(t, byte('x'), slice[0])

}

// ## Append
// The `append(<SLICE>, <VALUE1>, <VALUE2>, ...)` function
// allows adding new elements to an existing slice. It is
// often said that `append()` allows to _grow_ a slice.
func Test_Append(t *testing.T) {

	var sliceInt []int
	var sliceStr []string

	var sliceIntPrime = append(sliceInt, 15)
	var sliceStrPrime = append(sliceStr, "Einstein", "Newton")

	// Assertions
	assert.Equal(t, 1, len(sliceIntPrime))        // New length
	assert.Equal(t, 2, len(sliceStrPrime))        // New length
	assert.Equal(t, 15, sliceIntPrime[0])         // New element #0
	assert.Equal(t, "Einstein", sliceStrPrime[0]) // New element #0
	assert.Equal(t, "Newton", sliceStrPrime[1])   // New element #1
}

// ## Copy
// The `copy(<DST_SLICE>, <SRC_SLICE>)` function allows copying
// the elements from one slice onto another one.
func Test_Copy(t *testing.T) {
	// ------------- Same length --------------
	src1 := []byte{'a', 'b', 'c'}
	dst1 := []byte{'x', 'y', 'z'}
	copy(dst1, src1)

	// Assertions
	assert.Equal(t, []byte{'a', 'b', 'c'}, dst1)
	assert.Equal(t, dst1, src1)

	// --- Source smaller than destination  ---
	src2 := []byte{'a', 'b'}
	dst2 := []byte{'x', 'y', 'z'}
	copy(dst2, src2)

	// Assertions
	assert.Equal(t, []byte{'a', 'b', 'z'}, dst2)
	assert.NotEqual(t, dst2, src2)

	// ---- Source larger than destination ----
	src3 := []byte{'a', 'b', 'c', 'd'}
	dst3 := []byte{'x', 'y', 'z'}
	copy(dst3, src3)

	// Assertions
	assert.Equal(t, []byte{'a', 'b', 'c'}, dst3)
	assert.NotEqual(t, dst3, src3)
}

// ## Subsets (Slices, Splits, Trims)
// There are various ways to obtain a subset of a slice; using
// the slice notation as well as helper functions.
func Test_Slice_Subsets(t *testing.T) {
	// Assertions (Slice Notation)
	assert.Equal(t, byte('e'), []byte("hello")[1])
	assert.Equal(t, []byte("ello"), []byte("hello")[1:])
	assert.Equal(t, []byte("hel"), []byte("hello")[:3])
	assert.Equal(t, []byte("ell"), []byte("hello")[1:4])
	assert.Equal(t, []byte("ell"), []byte("hello")[1:4])

	// Assertions (Helper Functions)
	assert.Equal(t, [][]byte{[]byte("a"), []byte("b"), []byte("c")}, bytes.Split([]byte("a b c"), []byte(" ")))
	assert.Equal(t, []byte("hello"), bytes.Trim([]byte("😊hello🍺"), "😊🍺"))
	assert.Equal(t, []byte("hello"), bytes.Trim([]byte("😊hello🍺"), "😊🍺"))
	assert.Equal(t, []byte("hello"), bytes.TrimSpace([]byte("\t \t hello\n")))
}

// ## Iteration (Classic)
// A slice may be traversed just like an array by
// counting elements until reaching the slice's length
// using a `for()` loop.
func Test_Traverse(t *testing.T) {

	slice := []byte{'a', 'b', 'c'}
	var result = ""
	var counter = 0
	for i := 0; i < len(slice); i++ {
		counter++
		result = result + string(slice[i])
	}
	assert.Equal(t, 3, counter)
	assert.Equal(t, "abc", result)
}

//
// ## Iteration (Range)
//
// See [Control Flow](#control-flow)
//
// ## Pass By Value vs Pass By Reference
// Slices are passed to functions by reference by default.
func Test_Reference(t *testing.T) {
	changeFirstElement := func(slice []byte) {
		slice[0] = 'y'
	}

	slice := []byte{'a', 'b', 'c'}

	// Assertion #1
	// A copy of a slice still points to the original slice
	slice2 := slice
	slice2[0] = 'x'
	assert.Equal(t, byte('x'), slice[0]) // Rather than a

	// Assertion #2
	// Functions receive an slice pointer, rather than a copy
	changeFirstElement(slice)
	assert.Equal(t, byte('y'), slice[0]) // Rather than a or x

	// Assertion #3
	// A slice's slice still points to the original slice
	lastTwoLetters := slice[1:]
	lastTwoLetters[0] = 'z'
	assert.Equal(t, byte('z'), slice[1]) // Rather than b

}
