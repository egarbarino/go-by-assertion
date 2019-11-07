// # Functions
// Ignore-On
package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Functions With Multiple Return Values
func Test_Multiple_Return(t *testing.T) {
	first, second := returnTwo()
	assert.Equal(t, 'a', first)
	assert.Equal(t, false, second)

	onlyFirst, _ := returnTwo()
	assert.Equal(t, 'a', onlyFirst)

	_, onlySecond := returnTwo()
	assert.Equal(t, false, onlySecond)

}

func returnTwo() (rune, bool) {
	return 'a', false
}

// ## Variadic Functions
func Test_Variadic(t *testing.T) {
	assert.Equal(t, 0, sum())
	assert.Equal(t, 1, sum(1))
	assert.Equal(t, 3, sum(1, 2))
	assert.Equal(t, 6, sum(1, 2, 3))
}

func sum(numbers ...int) int {
	var result int = 0
	for _, n := range numbers {
		result = result + n
	}
	return result
}

// ## Anonymous Functions
func Test_Anonymous(t *testing.T) {

	// Direct use
	assert.Equal(t, 3, operation("sum")(1, 2))
	assert.Equal(t, 6, operation("mul")(2, 3))

	// Function extraction first
	sum := operation("sum")
	mul := operation("mul")
	assert.Equal(t, 3, sum(1, 2))
	assert.Equal(t, 6, mul(2, 3))

	// Create Inline
	assert.Equal(t, 3, func(v1 int, v2 int) int { return v1 + v2 }(1, 2))

}

func operation(op string) func(int, int) int {
	switch op {
	case "sum":
		return func(v1 int, v2 int) int {
			return v1 + v2
		}
	case "mul":
		return func(v1 int, v2 int) int {
			return v1 * v2
		}
	default:
		panic("op should be either \"sum\" or \"mul\"")
	}
}

// ## Closures
func Test_Closures(t *testing.T) {
	var counter = 0
	incrementer := func() int {
		counter++
		return counter
	}
	assert.Equal(t, 1, incrementer())
	counter++
	assert.Equal(t, 3, incrementer())
	incrementer()
	assert.Equal(t, 5, incrementer())
}
