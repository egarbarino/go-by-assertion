// # Strings
// In Go, a string is a discrete type as opposed
// to a list, or an array, like in Haskell, or C,
// respectively. However, it behaves like a slice
// of bytes in most cases.
//
// Ignore-On
package strings

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Length
// A string's length is obtained using the `len()` function.
func Test_String_Length(t *testing.T) {
	// Assertions
	assert.Equal(t, 0, len(""))
	assert.Equal(t, 5, len("hello"))
}

// ## Substring and Rune Index
// The following family of functions from the package
// `strings` provide the location of strings and characters
// (runes) in various contexts.
func Test_String_Index(t *testing.T) {
	// Assertions
	assert.Equal(t, 1, strings.Index("hello fellows", "ell"))
	assert.Equal(t, 7, strings.LastIndex("hello fellows", "ell"))
	assert.Equal(t, 2, strings.IndexByte("hello", 'l'))
	assert.Equal(t, 3, strings.LastIndexByte("hello", 'l'))
	assert.Equal(t, 4, strings.IndexRune("🍷😊🍺🍷", '😊'))
	assert.Equal(t, 2, strings.IndexAny("hi😊!", "😊🍺🍷"))
	assert.Equal(t, 6, strings.IndexAny("🤣hi😊!", "😊🍺🍷"))
	assert.Equal(t, 12, strings.LastIndexAny("🍷😊🍺🍷", "🍷"))
}

// ## Predicates (Substring, Starts/Ends With)
// The following family of functions from the package
// `strings` provide assertions about various substring
// membership cases.
func Test_String_Predicates(t *testing.T) {
	// Assertions
	assert.Equal(t, true, strings.Contains("hello", "llo"))
	assert.Equal(t, true, strings.ContainsAny("hello", "elxqwer"))
	assert.Equal(t, true, strings.ContainsAny("hi😊!", "😊🍺🍷"))
	assert.Equal(t, true, strings.HasPrefix("hello", "he"))
	assert.Equal(t, true, strings.HasSuffix("hello", "lo"))
}

// ## Subsets (Slices, Splits, Trims)
// The slice syntax `[start:end]` and the below functions
//  from the package `strings` provide subsets of a string.
func Test_String_Subsets(t *testing.T) {
	// Assertions
	assert.Equal(t, "e", string("hello"[1]))
	assert.Equal(t, "ello", "hello"[1:])
	assert.Equal(t, "hel", "hello"[:3])
	assert.Equal(t, "ell", "hello"[1:4])
	assert.Equal(t, "ell", "hello"[1:4])
	assert.Equal(t, []string{"a", "b", "c"}, strings.Split("a b c", " "))
	assert.Equal(t, "hello", strings.Trim("😊hello🍺", "😊🍺"))
	assert.Equal(t, "hello", strings.Trim("😊hello🍺", "😊🍺"))
	assert.Equal(t, "hello", strings.TrimSpace("\t \t hello\n"))
}

// ## Replace
// The `Replace()` function from the `strings` package
// substitutes one string with another.
func Test_String_Replace(t *testing.T) {
	// Assertions
	assert.Equal(t, "hello 😊😊", strings.Replace("hello :):)", ":)", "😊", -1))
	assert.Equal(t, "hello 😊:)", strings.Replace("hello :):)", ":)", "😊", 1))
}

// ## Case Conversion (Title, Upper, Lower)
// The `Title()`, `ToUpper()`, and `ToLower()` functions
// from the `strings` package change a string's case.
func Test_String_Convert_Case(t *testing.T) {
	// Assertions
	assert.Equal(t, "This Is The End", strings.Title("this is the end"))
	assert.Equal(t, "HELLO", strings.ToUpper("HellO"))
	assert.Equal(t, "hello", strings.ToLower("HellO"))
}

// ## Join (Concat)
// The `Join()` function from the `strings` package
// allows concatenating strings.
func Test_String_Join(t *testing.T) {
	// Assertions
	assert.Equal(t, "Hello-world", strings.Join([]string{"Hello", "world"}, "-"))
}

// ## Repeat
// The `Repeat()` function from the `strings` package
// creates a string by repeating another one a set
// number of times.
func Test_String_Repeat(t *testing.T) {
	// Assertions
	assert.Equal(t, "hihihi", strings.Repeat("hi", 3))
}

// ## String Formatting
// The `Sprtinf()` function from the `fmt` package
// helps formatting various value types as well as
// adding padding (spaces) for alignment purposes.
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
