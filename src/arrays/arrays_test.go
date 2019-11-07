// Ignore-On
package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// # Arrays

// ## Initialisation
func Test_Init(t *testing.T) {

	var arrayEmpty1 [0]int
	arrayEmpty2 := [...]int{}

	assert.Equal(t, arrayEmpty1, arrayEmpty2)
	assert.Equal(t, 0, len(arrayEmpty1))

}

// ## Array Literals
func Test_Literals(t *testing.T) {

	// Explicit length
	arrayTwo := [2]string{"Einstein", "Newton"}
	// No need to specify length
	arrayThree := [...]string{"Einstein", "Newton", "Curie"}

	assert.Equal(t, 2, len(arrayTwo))
	assert.Equal(t, 3, len(arrayThree))
}

// ## Set Element's Values
func Test_SetValues(t *testing.T) {

	var arrayOne [1]int
	var arrayTwo [2]string

	arrayTwo[0] = "Einstein"
	arrayTwo[1] = "Newton"

	assert.Equal(t, 1, len(arrayOne))
	assert.Equal(t, "Einstein", arrayTwo[0])
	assert.Equal(t, "Newton", arrayTwo[1])
}

// ## Traversal
func Test_Traverse(t *testing.T) {

	array := [3]byte{'a', 'b', 'c'}
	var result = ""
	var counter = 0
	for i := 0; i < len(array); i++ {
		counter++
		result = result + string(array[i])
	}
	assert.Equal(t, 3, counter)
	assert.Equal(t, "abc", result)
}

// ## Pass By Value vs Pass By Reference
func Test_Reference(t *testing.T) {

	array := [3]byte{'a', 'b', 'c'}

	// Copying an array produces a brand new one
	array2 := array
	array2[0] = 'x'
	assert.Equal(t, byte('a'), array[0]) // Rather than x

	// Functions receive an array copy, rather than a reference
	changeFirstElement(array)
	assert.Equal(t, byte('a'), array[0]) // Rather than y

}

func changeFirstElement(array [3]byte) {
	array[0] = 'y'
}
