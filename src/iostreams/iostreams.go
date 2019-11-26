// Ignore-On

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ## Interacting with Stdin, Stdout, and Stderr
// The `os` package provides the `Stdin`, `Stdout`, and `Stderr`
// streams. The below example reads all data from `Stdin`,
// converts said data to upper case, and then writes the results
// to `Stdout`. It then uses the `Stderr` stream to report the
// number of bytes read.
func main() {
	// Read all Stdin
	data, _ := ioutil.ReadAll(os.Stdin)
	// Convert data to upper case and write to Stdout
	count, _ := os.Stdout.WriteString(strings.ToUpper(string(data)))
	// Write the number of bytes read to Stderr
	os.Stderr.WriteString(fmt.Sprintf("Bytes read: %d", count))
}
