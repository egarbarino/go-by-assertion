// Ignore-On
package maps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## String Formatting
func Test_String_Formatting(t *testing.T) {
	assert.Equal(t, "17", fmt.Sprintf("%d", 17))             // Decimal
	assert.Equal(t, "11", fmt.Sprintf("%x", 17))             // Hexadecimal
	assert.Equal(t, "21", fmt.Sprintf("%o", 17))             // Octal
	assert.Equal(t, "10001", fmt.Sprintf("%b", 17))          // Binary
	assert.Equal(t, "10001", fmt.Sprintf("%b", 17))          // Binary
	assert.Equal(t, "1.750000", fmt.Sprintf("%f", 1.75))     // Floating Point
	assert.Equal(t, "1.75", fmt.Sprintf("%.2f", 1.75))       // Floating Point
	assert.Equal(t, "true", fmt.Sprintf("%t", true))         // Boolean
	assert.Equal(t, "😊", fmt.Sprintf("%c", '😊'))             // Rune (Unicode point)
	assert.Equal(t, "hello", fmt.Sprintf("%s", "hello"))     // Rune (Unicode point)
	assert.Equal(t, "    o", fmt.Sprintf("%5s", "o"))        // Right Alignment
	assert.Equal(t, "h    ", fmt.Sprintf("%-5s", "h"))       // Left Alignment
	assert.Equal(t, "'😊'", fmt.Sprintf("%q", '😊'))           // Quoted Rune
	assert.Equal(t, "'😊'", fmt.Sprintf("%q", '😊'))           // Quoted Rune
	assert.Equal(t, "\"hello\"", fmt.Sprintf("%q", "hello")) // Quoted String
	assert.Equal(t, "c64", fmt.Sprintf("%v%v", "c", 64))     // Default String formatting
	assert.Equal(t, "int", fmt.Sprintf("%T", 64))            // Inferred Value's Type
	assert.Equal(t, "%", fmt.Sprintf("%%"))                  // Literal Percent Sign
}
