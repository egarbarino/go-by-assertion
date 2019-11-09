// # Go Routines
// Ignore-On
package strings

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ## Simple Goroutine Example
// We first declare a function that creates an artificial pause and that
// also updates a status slice of bool values to inform that it has
// completed:
func wasteTime(sleepSeconds int, status []bool, identifier int) {
	time.Sleep(time.Duration(sleepSeconds) * time.Second)
	status[identifier] = true
}

// We then call the above function using the prefix `go` which will
// cause it to run concurrently.
func Test_GoRoutines(t *testing.T) {
	status := []bool{false, false, false}

	go wasteTime(3, status, 0)
	go wasteTime(5, status, 1)
	go wasteTime(7, status, 2)

	// Assertions
	assert.Equal(t, []bool{false, false, false}, status) // All sleeping yet
	time.Sleep(4 * time.Second)                          // Wait 4 seconds
	assert.Equal(t, []bool{true, false, false}, status)  // wasteTime() 1 ran!
	time.Sleep(2 * time.Second)                          // Wait 2 seconds
	assert.Equal(t, []bool{true, true, false}, status)   // wasteTime() 2 ran!
	time.Sleep(2 * time.Second)                          // Wait 2 seconds
	assert.Equal(t, []bool{true, true, true}, status)    // wasteTime() 3 ran!
}
