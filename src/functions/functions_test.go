// # Functions
// Go functions are actually "procedures" rather than pure functions
// like in a FP language such as Haskell. It means that they are not
// necessarily idempotent; they may have side effects.
//
// Ignore-On
package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ignore-Off
// ## Simple Function
// A simple function neither takes nor returns values.
var globalVar = ""

func simple() {
	globalVar = "changed"
}

func Test_Simple(t *testing.T) {
	// Assertions
	assert.Equal(t, "", globalVar)
	simple()
	assert.Equal(t, "changed", globalVar)
}

// ## Simple Function With Return Value
// The return value type is declared after the parenthesis:
// `func name() <TYPE>`. The value itself is returned using
// `return <VALUE>`.
func hello() string {
	return "hello"
}

func Test_Simple_Return(t *testing.T) {
	// Assertions
	assert.Equal(t, "hello", hello())
}

// ## Function With Multiple Return Values
// Two or more values may be returned by a function; their type must be declared using
// the `func name() (TYPE1, TYPE2, ...)` syntax. When invoking the function, there is
// the flexibility to assign only the required values as shown below:
func returnTwo() (rune, bool) {
	return 'a', false
}

func Test_Multiple_Return(t *testing.T) {
	// Assertions (Two Values)
	first, second := returnTwo()
	assert.Equal(t, 'a', first)
	assert.Equal(t, false, second)

	// Assertions (First Value Only)
	onlyFirst, _ := returnTwo()
	assert.Equal(t, 'a', onlyFirst)

	// Assertions (Second Value Only)
	_, onlySecond := returnTwo()
	assert.Equal(t, false, onlySecond)

}

// ## Function Arguments
// Functions arguments are expressed using the `func name(<VAR> <TYPE>)`
// syntax. Multiple arguments may be declared by separating them with a comma
// as shown below:
var result = ""

func oneArgument(name string) {
	result = name
}

func twoArguments(name string, age int) {
	result = fmt.Sprintf("%s-%d", name, age)
}

func Test_Arguments(t *testing.T) {
	// Assertions (One Argument)
	oneArgument("Rob")
	assert.Equal(t, "Rob", result)

	// Assertions (Two Arguments)
	twoArguments("Rob", 36)
	assert.Equal(t, "Rob-36", result)

}

// ## Variadic Functions
// Variadic functions take a varying number of arguments. The syntax
// is `func name(<VAR> ...<TYPE>)`. Please note that `<TYPE>` is
// prefixed with three dots `...`. In the function's body, `<VAR>` will
// appear like an array and can be iterated using the `range` keyword.
func sum(numbers ...int) int {
	var result int = 0
	for _, n := range numbers {
		result = result + n
	}
	return result
}

func Test_Variadic(t *testing.T) {
	// Assertions
	assert.Equal(t, 0, sum())
	assert.Equal(t, 1, sum(1))
	assert.Equal(t, 3, sum(1, 2))
	assert.Equal(t, 6, sum(1, 2, 3))
}

// ## Anonymous Functions
// Anonymous functions are declared without a custom name using the form `func() {...}`
// and can be assigned to variables or returned directly as values.
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

func Test_Anonymous(t *testing.T) {
	// Assertions

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

// ## Closures
// Closures are a type of function that references variables from
// outside its body.
func Test_Closures(t *testing.T) {
	var counter = 0

	incrementer := func() int {
		counter++
		return counter
	}

	// Assertions
	assert.Equal(t, 1, incrementer())
	counter++
	assert.Equal(t, 3, incrementer())
	incrementer()
	assert.Equal(t, 5, incrementer())
	assert.Equal(t, 6, incrementer())
}
