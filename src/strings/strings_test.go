// # Strings
// Ignore-On
package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Quantifiers (Length, Occurrances)
func Test_String_Length(t *testing.T) {
	assert.Equal(t, 0, len(""))
	assert.Equal(t, 5, len("hello"))
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
func Test_String_Predicates(t *testing.T) {
	assert.Equal(t, true, strings.Contains("hello", "llo"))
	assert.Equal(t, true, strings.ContainsAny("hello", "elxqwer"))
	assert.Equal(t, true, strings.ContainsAny("hi😊!", "😊🍺🍷"))
	assert.Equal(t, true, strings.HasPrefix("hello", "he"))
	assert.Equal(t, true, strings.HasSuffix("hello", "lo"))
}

// ## Subsets (Slices, Splits, Trims)
func Test_String_Subsets(t *testing.T) {
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
func Test_String_Replace(t *testing.T) {
	assert.Equal(t, "hello 😊😊", strings.Replace("hello :):)", ":)", "😊", -1))
	assert.Equal(t, "hello 😊:)", strings.Replace("hello :):)", ":)", "😊", 1))
}

// ## Convert Case (Title, Upper, Lower)
func Test_String_Convert_Case(t *testing.T) {
	assert.Equal(t, "This Is The End", strings.Title("this is the end"))
	assert.Equal(t, "HELLO", strings.ToUpper("HellO"))
	assert.Equal(t, "hello", strings.ToLower("HellO"))
}

// ## Join (Concat) and Repeat
func Test_String_Generation(t *testing.T) {
	assert.Equal(t, "Hello-world", strings.Join([]string{"Hello", "world"}, "-"))
	assert.Equal(t, "hihihi", strings.Repeat("hi", 3))
}
