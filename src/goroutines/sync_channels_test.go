// Ignore-On
package strings

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ## Synchronous Channels
// Synchronous channels are first-in-first-out (FIFO) queues.
// A channel is created using `make(chan <TYPE>)`:

// Here we poll the queue using `<-foodChannel`. Messages may not arrive
// in sequence but we have added a sleep statement to `platter()` so that
// they do.
func Test_Channel(t *testing.T) {

	// Create foodChannel queue
	var foodChannel = make(chan string)

	// The below function adds one or more food items to the `foodChannel` queue.
	platter := func(size int) {
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
// that their "blocking" behaviour can be appreciated.
func Test_Channel2(t *testing.T) {
	// Define a channel c
	c := make(chan string)

	execOrder := "A"

	// Run a goroutine that posts "message" to c
	go func() {

		execOrder = execOrder + "B"
		c <- "message" // It blocks until a receiver appears
		execOrder = execOrder + "E"

	}()

	// Wait
	time.Sleep(time.Second)

	// Poll queue: receive message
	execOrder = execOrder + "C"
	message := <-c
	execOrder = execOrder + "D"

	time.Sleep(time.Second) // Sleep 1 second

	// Assertions
	assert.Equal(t, "message", message)
	assert.Equal(t, "ABCDE", execOrder)
}

// ## Channel Buffering
// By default, channels block the sender until a receiver collects the
// message. With buffering, messages can be "held" in the queue and
// extracted later, decoupling the sender from the receiver. This is
// accomplished by specifying the number of messages to be buffered as
// second argument (N) to `make(chan <TYPE>, N)`.
func Test_Channel3(t *testing.T) {

	c := make(chan string, 2)
	c <- "first"  // Non-blocking
	c <- "second" // Non-blocking
	// c <- "third"  // This one would block

	// Assertions
	assert.Equal(t, "first", <-c)
	assert.Equal(t, "second", <-c)

}

// ## Channel Directions
// Functions can specify whether channels provided as arguments can
// only send, receive, (or both send a receive).
// This is accomplished by specifying the channel signature as either
// `<-chan <TYPE>`, `chan<-<TYPE>`, or simply `chan <TYPE>`, respectively.
func Test_Channel5(t *testing.T) {
	c := make(chan string, 1)

	canOnlySend := func(c chan<- string, msg string) {
		c <- msg
		// msg := <-c // Won't compile
	}
	canOnlyReceive := func(c <-chan string) string {
		// c <- msg // Won't compile
		return (<-c)
	}
	canBothSendAndReceive := func(c chan string, msg string) string {
		c <- msg
		return (<-c)
	}

	// Assertions
	canOnlySend(c, "hello")
	assert.Equal(t, "hello", canOnlyReceive(c))
	assert.Equal(t, "hola!", canBothSendAndReceive(c, "hola!"))
}

// ## Select
// The `select { ...}` statement is similar to the `switch` statement,
// but it is used to wait for messages in multiple channels simultaneously.

// In the below example, the message _ping_ is passed to channel `c1`, `c2`,
// and finally `c3` in each iteration of the `for` loop:
func Test_Channel7(t *testing.T) {
	execOrder := "A"

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)

	c1 <- "ping"

	execOrder = execOrder + "B"

	for i := 0; i < 3; i++ { // Loop three times
		execOrder = execOrder + "-"
		select {
		case msg := <-c1: // Assignment is optional
			execOrder = execOrder + "c1(" + msg + ")"
			c2 <- msg
		case msg := <-c2:
			execOrder = execOrder + "c2(" + msg + ")"
			c3 <- msg
		case msg := <-c3:
			execOrder = execOrder + "c3(" + msg + ")"
		}
	}

	// Assertions
	assert.Equal(t, "AB-c1(ping)-c2(ping)-c3(ping)", execOrder)
}

// ## Select and Timeouts
// A `select {...}` statement will block until a message is received
// by one of the cases declared in it. It is possible to frame one of said
// cases as a timeout condition using the `case: <- <Time>` signature.

// In the below example, the timeout case is hit in the first iteration,
// whereas the message receipt is hit is  in the second one.
func Test_Channel9(t *testing.T) {
	execOrder := "A"

	c := make(chan string, 1)

	for i := 0; i < 2; i++ { // Loop two times
		select {
		case <-c: // wait for any message
			execOrder = execOrder + "C"
		case <-time.After(1 * time.Second):
			execOrder = execOrder + "B"
		}
		c <- "ping"
	}

	// Assertions
	assert.Equal(t, "ABC", execOrder)
}

// ## Select with "Send" Cases
// The `select {...}` statement can also be used with send cases in
// the form `case channel <- msg`. This is helpful to avoid a blocking
// situation whenever there are no receivers associated with the channel.

// In the below example, the timeout case is triggered because `c`
// is a synchronous channel and there are no receivers waiting on it.
func Test_Channel12(t *testing.T) {
	execOrder := "A"

	c := make(chan string) // synchronous channel!

	select {
	case c <- "ping": // would block because there is no receiver
		execOrder = execOrder + "Never"
	case <-time.After(1 * time.Second): // therefore the result is a timeout
		execOrder = execOrder + "B"
	}

	// Assertions
	assert.Equal(t, "AB", execOrder)
}

// ## Select with Defaults
// Sometimes we want to tell whether there is a message pending
// to be retrieved (or a receiver for a channel to which we are sending a
// message) in an immediate manner.
// In this case, there is no need to create an artificial timeout,
// the keyword `default` can be used instead.

// In the below example, the default case is triggered because `c`
// is a synchronous channel and there are no receivers waiting on it.
func Test_Channel15(t *testing.T) {
	execOrder := "A"

	c := make(chan string) // synchronous channel!

	select {
	case c <- "ping": // would block because there is no receiver
		execOrder = execOrder + "Never"
	default: // therefore the result is the default case
		execOrder = execOrder + "B"
	}

	// Assertions
	assert.Equal(t, "AB", execOrder)
}

// ## Opening and Closing Channels
// Channels are open by default when created. They can explicitly be closed
// using the `close(channel)` statement.
// The status of a channel may be checked using the second return value
// of the message receive statement: `msg, channelStatus <- channel`.
// Channels can only be closed if they contain no pending messages
// in the buffer. Furthermore, channels can be closed only once.

// In the below example we iterate two times through a declaration
// that obtains both the message and the status, and that closes the
// channel upon receiving the first message.
func Test_Channel17(t *testing.T) {
	execOrder := "A"

	c := make(chan string, 3)
	c <- "msg1"

	for i := 0; i < 2; i++ { // loop two times
		msg, channelStillOpen := <-c
		if channelStillOpen {
			execOrder = execOrder + "B(" + msg + ")"
			close(c)
		} else {
			execOrder = execOrder + "C(" + msg + ")"
		}
	}

	// Assertions
	assert.Equal(t, "AB(msg1)C()", execOrder)
}

// ## Range over Channels
// It is possible to iterate over the received messages on a channel
// using the `range` keyword.
func Test_Channel19(t *testing.T) {
	execOrder := "A-"
	c := make(chan string, 3)
	c <- "m1"
	c <- "m2"
	c <- "m3"
	close(c) // range will remain suspended unless we close the channel
	for msg := range c {
		execOrder = execOrder + msg
	}
	assert.Equal(t, "A-m1m2m3", execOrder)
}

// ## Scheduled Tasks (Timers)
// A timer, created via `timer.NewTimer(Time)` offers a mechanism to
// run goroutines in the future. The implementation, as appreciated
// below, relies on a blocking channel that forces the goroutine to wait
// until the specified time (5 seconds).
func Test_Channel25(t *testing.T) {
	execOrder := "A"

	timer := time.NewTimer(5 * time.Second)

	futureFunc := func() {
		execOrder = execOrder + "B"
		<-timer.C // This is what makes the function wait!
		execOrder = execOrder + "D"
	}

	go futureFunc()
	time.Sleep(2 * time.Second)

	execOrder = execOrder + "C"

	time.Sleep(4 * time.Second) // 6 seconds elapsed by now.

	execOrder = execOrder + "E"

	// Assertions
	assert.Equal(t, "ABCDE", execOrder)
}

// ## Scheduled Tasks (Cancelling Timers)
// A timer may be aborted by invoking `timer.Stop()`, preventing the
//  waiting goroutine from executing
// the code after the channel _receive_ statement.
func Test_Channel30(t *testing.T) {
	execOrder := "A"

	timer := time.NewTimer(5 * time.Second)

	futureFunc := func() {
		execOrder = execOrder + "B"
		<-timer.C
		execOrder = execOrder + "Never" // will not be reached
	}

	go futureFunc()
	time.Sleep(2 * time.Second)

	execOrder = execOrder + "C"

	timer.Stop() // Abort timer!

	execOrder = execOrder + "D"

	// Assertions
	assert.Equal(t, "ABCD", execOrder)
}

// ## Recurrent Tasks (Tickers)
// Recurrent tasks may be set up using `time.NewTicker(time)` and applied
// similarly to regular timers, except that the recurrent code must run
// inside some form of repeating loop as shown below.

// Please note that the goroutine in the example terminates because
// the main function ends as well. It may be necessary,
// in a real world scenario, to implement
// a mechanism (i.e. a sentinel message via a channel) to force the
// repeating loop to exit.
func Test_Channel40(t *testing.T) {
	execOrder := "A"

	timer := time.NewTicker(time.Second) // 1 Second

	repeatingFunc := func() {
		for {
			execOrder = execOrder + "<"
			<-timer.C
			execOrder = execOrder + ">"
		}
	}

	go repeatingFunc()

	time.Sleep(3500 * time.Millisecond) // 3.5 seconds
	timer.Stop()                        // Abort timer!

	execOrder = execOrder + "B"

	// Assertions
	assert.Equal(t, "A<><><><B", execOrder)
}

// ## Race Conditions (Mutexes)
// Concurrent goroutines that share a common variable
// may result in inconsistent behaviour whenever operations again
// said variable aren't atomic.

// In the below example, we execute 10,000 times both the `increment()`
// and `decrement()` goroutines, which should result in `x == 0` at the
// end, but this is seldom the case.
func Test_Channel50(t *testing.T) {
	var x = 0

	var increment = func() {
		x = x + 1
	}
	var decrement = func() {
		x = x - 1
	}

	for i := 0; i < 10000; i++ {
		go increment()
		go decrement()
	}

	// Assertions
	assert.NotEqual(t, 0, x) // There is a tiny chance this may be true :)
}

// The reason why is that `x = x - 1` is not an atomic operation; it may
// involve multiple steps such as the ones shown in the below pseudocode:

// ```
// temp = get value of X
// temp2 = 1 + temp2
// x = temp2
// ```

// The solution is to enclose the `x = x + 1` statement with `mutex.Lock()`
// and `mutex.Unlock()` methods from `sync.Mutex` so that only one goroutine
// has access to either `x = x + 1` or `x = x - 1` at any given time:

func Test_Channel60(t *testing.T) {
	var mutex sync.Mutex
	var x = 0

	var increment = func() {
		mutex.Lock()
		x = x + 1
		mutex.Unlock()
	}
	var decrement = func() {
		mutex.Lock()
		x = x - 1
		mutex.Unlock()
	}

	for i := 0; i < 10000; i++ {
		go increment()
		go decrement()
	}

	time.Sleep(5 * time.Second) // Sleep 5 seconds

	// Assertions
	assert.Equal(t, 0, x)
}

// ## Race Conditions (Atomic Counters)
// In the case of simple numerical counters, there is no need to set up
// mutexes, or channels. The `atomic.Add<TYPE>(&counter, <N>)` function
// allows to alter a counter in a thread-safe manner as shown in the
// example below. There are also other functions such as
// `Store<TYPE>(&counter, <N>)` which
// simply sets a new value rather than adding to an existing one.
func Test_Channel55(t *testing.T) {
	var x int32 = 0
	var increment = func() {
		atomic.AddInt32(&x, 1)
	}
	var decrement = func() {
		atomic.AddInt32(&x, -1)
	}

	for i := 0; i < 10000; i++ {
		go increment()
		go decrement()
	}
	time.Sleep(3 * time.Second)

	// Assertions
	assert.Equal(t, int32(0), atomic.LoadInt32(&x))
	assert.Equal(t, int32(0), x) // May be inconsistent if goroutines are still active
}

// ## Wait Groups
// Wait Groups allow to treat a swarm of concurrent goroutines as one
// logical unit of execution to spare the programmer from coordinating
// the orchestration of each individual goroutine.

// The `sync.WaitGroup` type is used for this purpose
// and uses the below methods:

// * `Add(N)` - tell how many goroutines are pending: `pending++`
// * `Done()` - tell that a goroutine has finished: `pending--`
// * `Wait()` - block flow until `pending == 0`

// In the below example, we calculate the two times table in parallel,
// using `sync.WaitGroup` to coordinate the task:
func Test_Channel80(t *testing.T) {

	var waitGroup sync.WaitGroup // Allocate a new Wait Group. pending := 0
	twoTimesTable := make([]int, 10)

	var double = func(number int, wg *sync.WaitGroup) {
		time.Sleep(3 * time.Second) // Artificial wait
		twoTimesTable[number] = (number + 1) * 2
		wg.Done() // Goroutine completed: pending--
	}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1) // Add new pending goroutine: pending++
		go double(i, &waitGroup)
	}

	waitGroup.Wait() // Wait until pending == 0

	// Assertions
	assert.Equal(t, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, twoTimesTable)
}
