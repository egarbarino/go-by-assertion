// Ignore-On
package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Interacting with Shell Commands
// It is possible to interact with commands by using the `Command()` function
// from the `exec` package. Other than passing arguments as separate function
// arguments to `Command()`, piping data in and out of commands involves
// connecting the `Stdin`, `Stdout`, and `Stderr` streams declared by the
// `Cmd` struct.
func Test_IOStreams(t *testing.T) {
	// Specify command
	cmd := exec.Command("go", "run", "iostreams.go")

	// Specify stdin input
	cmd.Stdin = strings.NewReader("hello")

	// Declare a buffer to capture stdout from command
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Declare a buffer to capture stderr from command
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run Command
	err := cmd.Run()
	if err != nil {
		t.Error("Error:", err)
	}

	// Assertions
	assert.Equal(t, "HELLO", stdout.String())
	assert.Equal(t, "Bytes read: 5", stderr.String())
}
