// # Go Routines
// Ignore-On
package strings

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ## Simple Goroutine
// A goroutine is nothing more than a regular function that is
// run asynchronously by prefixing it with the `go` keyboard.
// For example: `go myfunc()`.

// In the below example, the `execOrder` strings captures
// the actual order of execution
func Test_GoRoutines(t *testing.T) {
	// This function waits for a number of seconds and then appends
	// an identifier to the `execOrder` string.
	wasteTime := func(sleepSeconds int, execOrder *string, id string) {
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
		*execOrder = *execOrder + id
	}

	execOrder := "A"

	go wasteTime(7, &execOrder, "E") // will run last
	go wasteTime(5, &execOrder, "D") // will run second
	go wasteTime(3, &execOrder, "C") // will run first

	execOrder = execOrder + "B"

	// Wait 10 seconds to allow all goroutines to run.
	time.Sleep(time.Duration(10) * time.Second)

	execOrder = execOrder + "F"

	// Assertions
	assert.Equal(t, "ABCDEF", execOrder)

}
