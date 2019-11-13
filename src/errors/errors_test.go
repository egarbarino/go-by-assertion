// # Errors
// Ignore-On
package main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## The Error Interface
// Built-in functions that produce errors typically implement
// the `error` interface which has a method with the signature `Error()`.
func Test_Error_Interface(t *testing.T) {
	// Open a non-existing file
	_, err := os.Open("non-existing-file.txt")

	// Assertions
	assert.Equal(t, "open non-existing-file.txt: no such file or directory", err.Error())
}

// ## Custom Errors
// Customs errors are created using the `errors.New()` function.
func safeDiv(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return a / b, nil
}
func Test_Error_Custom(t *testing.T) {

	// Calculations
	r, _ := safeDiv(12, 3)
	_, err := safeDiv(12, 0)

	// Assertions
	assert.Equal(t, 4, r)
	assert.Equal(t, "Can't divide by zero", err.Error())
}

// ## Wrapping Errors with String Formatting
func Test_Error_Formatting(t *testing.T) {
	// Open a non-existing file
	_, err := os.Open("non-existing-file.txt")

	// Annotate error
	err2 := fmt.Errorf("Ouch! %s", err)

	// Assertions
	assert.Equal(t, "Ouch! open non-existing-file.txt: no such file or directory", err2.Error())
}

// ## Custom Error Struct
// A custom struct type can be used to provide more details about a given error.
// In the example below, `CustomError` can provide details about the
// line number and row number where an error occured.
type CustomError struct {
	LineNumber int
	RowNumber  int
}

// We then implement the `Error()` method to comply with the `error` interface:
func (c CustomError) Error() string {
	return fmt.Sprintf("Failed at line %d, row %d", c.LineNumber, c.RowNumber)
}

// The idea is that all errors abide to the `error` interface. To test this
// principle, we define below a function which may fail either due
// to a system error from `os.Open()` or one using our custom
// `CustomError` type:
func doSomethingNasty(option int) error {
	if option == 1 {
		_, err := os.Open("non-existing-file.txt")
		return err
	}
	return &CustomError{LineNumber: 5, RowNumber: 30}
}

// We finally test the two different errors using the type
// assertion `err.(*CustomerError)`.
func Test_Error_Struct(t *testing.T) {

	var err error
	var line, row int
	var path string

	// Trigger both doSomethingNasty(1) and doSomethingNasty(2)
	for i := 1; i <= 2; i++ {
		err = doSomethingNasty(i)
		switch err.(type) {
		case *os.PathError:
			v, _ := err.(*os.PathError)
			path = v.Path
		case *CustomError:
			v, _ := err.(*CustomError)
			line = v.LineNumber
			row = v.RowNumber
		}
	}

	// Assertions
	assert.Equal(t, "non-existing-file.txt", path)
	assert.Equal(t, 5, line)
	assert.Equal(t, 30, row)

}
