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
	// Verify stdout and stderr values
	assert.Equal(t, "HELLO", stdout.String())
	assert.Equal(t, "Bytes read: 5", stderr.String())
}
