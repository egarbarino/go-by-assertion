package flowcontrol

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## If-Then-Else
func Test_If(t *testing.T) {
	var successCounter = 0
	var failureCounter = 0
	// Simple If
	if 2*3 == 6 {
		successCounter++
	}
	// If Else
	if 2*3 == 7 {
		failureCounter++
	} else {
		successCounter++
	}
	// If Else If
	if 2*3 == 8 {
		failureCounter++
	} else if 2*3 == 7 {
		failureCounter++
	} else if 2*3 == 6 {
		successCounter++
	} else {
		failureCounter++
	}
	assert.Equal(t, 3, successCounter)
	assert.Equal(t, 0, failureCounter)
}

// ## Switch (Normal)
func Test_Switch_Normal(t *testing.T) {
	var result = ""
	switch 1 + 1 {
	case 0:
		result = "zero"
	case 1:
		result = "one"
	case 2:
		result = "two"
	default: // Optional
		result = "some_number"
	}
	assert.Equal(t, "two", result)
}

// ## Switch (Multi-Value Case)
func Test_Switch_Multi_Value(t *testing.T) {
	// Just cases
	var result = ""
	switch 2 * 3 {
	case 0, 1, 2, 3, 4, 5, 6:
		result = "between_zero_and_six"
	case 7:
		result = "seven"
	default: // Optional
		result = "some_number"
	}
	assert.Equal(t, "between_zero_and_six", result)
}

// ## Switch (Expression Cases)
func Test_Switch_Expression_Cases(t *testing.T) {
	var result = ""
	switch { // No expression here
	case 2*3 == 1:
		result = "one"
	case 2*3 == 6:
		result = "six"
	default: // Optional
		result = "some_number"
	}
	assert.Equal(t, "six", result)
}

// ## Switch (Type Cases)
func Test_Switch_Type_Cases(t *testing.T) {
	testType := func(i interface{}) string {
		switch t := i.(type) { // the t := assignment is optional
		case int:
			return "int"
		case string:
			return "string"
		default: // Optional
			return "other type: " + fmt.Sprintf("%T", t)
		}
	}
	assert.Equal(t, "string", testType("hello"))
	assert.Equal(t, "int", testType(45))
	assert.Equal(t, "other type: float64", testType(1.53))
}

// ## Inifinite Loop (While True Loop)
func Test_While_True_Loop(t *testing.T) {
	var counter = 0
	for {
		counter = counter + 1
		if counter == 3 {
			break
		}
	}
	assert.Equal(t, 3, counter)
}

// ## While Loop
func Test_While_Loop(t *testing.T) {
	var counter = 0
	for counter < 3 {
		counter = counter + 1
	}
	assert.Equal(t, 3, counter)
}

// ## Iterate Over Array Elements
func Test_Range_Array_Loop(t *testing.T) {
	var indexSum = 0
	var sum = 0
	for index, currentValue := range []int{1, 2, 3} {
		indexSum = indexSum + index // 0 + 1 + 2 = 3
		sum = sum + currentValue    // 1 + 2 + 3 = 6
	}
	assert.Equal(t, 3, indexSum)
	assert.Equal(t, 6, sum)
}

// ## Iterate Over Array Elements But Use Only Index
func Test_Range_Array_Loop_Ignore_Index(t *testing.T) {
	var indexSum = 0
	for index := range []int{50, 999, 300} {
		indexSum = indexSum + index // 0 + 1 + 2 = 3
	}
	assert.Equal(t, 3, indexSum)
}

// ## Iterate Over Array Elements But Ignore Index
func Test_Range_Array_Loop_Ignore_Index_Blank_Identifier(t *testing.T) {
	var sum = 0
	for _, currentValue := range []int{1, 2, 3} {
		sum = sum + currentValue // 1 + 2 + 3 = 6
	}
	assert.Equal(t, 6, sum)
}

// ## Iterate Over Keys and Values of a Map
func Test_Range_Map_Loop(t *testing.T) {
	var keys, values string
	for k, v := range map[string]string{"A": "(Argentina)", "B": "(Brazil)"} {
		keys = keys + k
		values = values + v
	}
	assert.Equal(t, "AB", keys)
	assert.Equal(t, "(Argentina)(Brazil)", values)
}

// Iterate Over the Unicode Characters of a String
func Test_Range_String_Loop(t *testing.T) {
	var word string
	for _, v := range "😊 olleh" {
		word = string(v) + word
	}
	assert.Equal(t, "hello 😊", word)
}

// Iterate Over The Bytes of a String
func Test_For_String_Bytes_Loop(t *testing.T) {
	var reversedHello = "😊 olleh"
	var word string
	for i := 0; i < len(reversedHello); i++ {
		word = string(reversedHello[i]) + word
	}
	assert.Equal(t, "hello \u008a\u0098\u009fð", word)
}

// ## Regular Foor Loop
func Test_Regular_For_Loop(t *testing.T) {
	var counter = 0
	for i := 0; i < 3; i++ {
		counter = counter + i // 0 + 1 + 2 = 3
	}
	assert.Equal(t, 3, counter)
}
