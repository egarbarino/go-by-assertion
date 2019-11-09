// Ignore-On
package strings

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ## Synchronous Channels
// Synchronous channels are first-in-first-out (FIFO) queues.
// A channel is created using `make(chan <TYPE>)`:

var foodChannel = make(chan string)

// The below function adds one or more food items to the `foodChannel` queue.
func platter(size int) {

	// Artificial wait
	time.Sleep(time.Duration(size) * time.Second)

	if size >= 1 {
		foodChannel <- "hummus" // Post "hummus" to the queue
	}
	if size >= 2 {
		foodChannel <- "falafel" // Post "falafel" to the queue
	}
	if size >= 3 {
		foodChannel <- fmt.Sprintf("%d pita", size-2) // Post "n pita" to the queue
	}
}

// Here we poll the queue using `<-foodChannel`. Messages may not arrive
// in sequence but we have added a sleep statement to `platter()` so that
// they do.
func Test_Channel(t *testing.T) {

	// Prepare various Platter Sizes
	go platter(1)
	go platter(2)
	go platter(3)

	// Assertions
	assert.Equal(t, "hummus", <-foodChannel)  // posted by platter(1)
	assert.Equal(t, "hummus", <-foodChannel)  // posted by platter(2)
	assert.Equal(t, "falafel", <-foodChannel) // posted by platter(2)
	assert.Equal(t, "hummus", <-foodChannel)  // posted by platter(3)
	assert.Equal(t, "falafel", <-foodChannel) // posted by platter(3)
	assert.Equal(t, "1 pita", <-foodChannel)  // posted by platter(3)
}

// ## Synchronous Channels (Blocking Behaviour)
// The following annotated code uses the `execOrder` string to
// demonstrate the order of execution using synchronous channels so
// that their "blocking" behaviour can be internalised.
func Test_Channel2(t *testing.T) {
	c := make(chan string)
	execOrder := "A"

	// Run a goroutine that posts "message" to c
	go func() {
		execOrder = execOrder + "B"
		c <- "message" // It blocks until a receiver appears
		execOrder = execOrder + "E"
	}()

	// Wait and poll queue
	time.Sleep(time.Second)
	execOrder = execOrder + "C"
	message := <-c
	execOrder = execOrder + "D"

	// Assertions
	assert.Equal(t, "message", message)
	assert.Equal(t, "ABCDE", execOrder)
}
