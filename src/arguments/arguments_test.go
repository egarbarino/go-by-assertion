// Ignore-On
package arguments

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Arguments
// A command's arguments are found in the `Args` slice from the `os` package.
// The first element is the command name whereas the second one, third and so
// on, are the arguments. In the example below, we check the command and
// arguments generated by the _test runner_.
func Test_Arguments(t *testing.T) {
	var arguments, command string
	for i := 0; i < len(os.Args); i++ {
		if i == 0 {
			command = os.Args[0]
		} else if i == 1 {
			arguments += os.Args[i]
		} else {
			arguments += " " + os.Args[i]
		}
	}

	// Assertions
	fmt.Printf(command)
	fmt.Printf(arguments)
	assert.Equal(t, true, strings.Contains(command, "arguments.test"))
	assert.Equal(t, true, strings.Contains(arguments, "test.timeout"))
}
