// Ignore-On
package slices

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// # Slices

// ## Initialisation
func Test_Init(t *testing.T) {

	var sliceEmpty1 []int
	var sliceEmpty2 = make([]int, 0, 0)
	var sliceEmpty3 = []int{}

	assert.Equal(t, sliceEmpty1, sliceEmpty2)
	assert.Equal(t, sliceEmpty2, sliceEmpty3)
	assert.Equal(t, 0, len(sliceEmpty1))
}

// ## Slice Literals
func Test_Literals(t *testing.T) {

	slice := []string{"Einstein", "Newton"}
	assert.Equal(t, 2, len(slice))

}

// ## Set Element's Values
func Test_Set_Values(t *testing.T) {

	slice := []byte{'a', 'b', 'c'}
	assert.Equal(t, byte('a'), slice[0])

	slice[0] = 'z'
	assert.Equal(t, byte('z'), slice[0])

}

// ## Append
func Test_Append(t *testing.T) {

	var sliceInt []int
	var sliceStr []string

	var sliceIntPrime = append(sliceInt, 15)
	var sliceStrPrime = append(sliceStr, "Einstein", "Newton")

	assert.Equal(t, 1, len(sliceIntPrime))
	assert.Equal(t, 2, len(sliceStrPrime))

	assert.Equal(t, 15, sliceIntPrime[0])
	assert.Equal(t, "Einstein", sliceStrPrime[0])
	assert.Equal(t, "Newton", sliceStrPrime[1])
}

// ## Append
func Test_Copy(t *testing.T) {

	// Same length
	src1 := []byte{'a', 'b', 'c'}
	dst1 := []byte{'x', 'y', 'z'}
	copy(dst1, src1)
	assert.Equal(t, []byte{'a', 'b', 'c'}, dst1)
	assert.Equal(t, dst1, src1)

	// Source smaller than destination
	src2 := []byte{'a', 'b'}
	dst2 := []byte{'x', 'y', 'z'}
	copy(dst2, src2)
	assert.Equal(t, []byte{'a', 'b', 'z'}, dst2)
	assert.NotEqual(t, dst2, src2)

	// Source larger than destination
	src3 := []byte{'a', 'b', 'c', 'd'}
	dst3 := []byte{'x', 'y', 'z'}
	copy(dst3, src3)
	assert.Equal(t, []byte{'a', 'b', 'c'}, dst3)
	assert.NotEqual(t, dst3, src3)
}

// ## Subsets (Slices, Splits, Trims)
func Test_Slice_Subsets(t *testing.T) {
	assert.Equal(t, byte('e'), []byte("hello")[1])
	assert.Equal(t, []byte("ello"), []byte("hello")[1:])
	assert.Equal(t, []byte("hel"), []byte("hello")[:3])
	assert.Equal(t, []byte("ell"), []byte("hello")[1:4])
	assert.Equal(t, []byte("ell"), []byte("hello")[1:4])
	assert.Equal(t, [][]byte{[]byte("a"), []byte("b"), []byte("c")}, bytes.Split([]byte("a b c"), []byte(" ")))
	assert.Equal(t, []byte("hello"), bytes.Trim([]byte("😊hello🍺"), "😊🍺"))
	assert.Equal(t, []byte("hello"), bytes.Trim([]byte("😊hello🍺"), "😊🍺"))
	assert.Equal(t, []byte("hello"), bytes.TrimSpace([]byte("\t \t hello\n")))
}

// ## Traversal
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

// ## Pass By Value vs Pass By Reference
func Test_Reference(t *testing.T) {

	slice := []byte{'a', 'b', 'c'}

	// A copy of a slice still points to the original slice
	slice2 := slice
	slice2[0] = 'x'
	assert.Equal(t, byte('x'), slice[0]) // Rather than a

	// Functions receive an slice pointer, rather than a copy
	changeFirstElement(slice)
	assert.Equal(t, byte('y'), slice[0]) // Rather than a or x

	// A slice's slice still points to the original slice
	lastTwoLetters := slice[1:]
	lastTwoLetters[0] = 'z'
	assert.Equal(t, byte('z'), slice[1]) // Rather than b

}

func changeFirstElement(slice []byte) {
	slice[0] = 'y'
}
