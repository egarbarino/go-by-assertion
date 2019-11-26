// # Maps
//
// Maps are associations of keys
// to values; both of which can be of
// any type.
// They are also known as _hashes_ or
// _hashmaps_ in other languages. Maps
// types have the form `map[<KEY-TYPE>]<VALUE-TYPE>`.
//
// Ignore-On
package maps

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Empty Map
// An empty map contains no key/value pairs.
func Test_Maps(t *testing.T) {
	var emptyMap1 = map[string]int{}
	var emptyMap2 = make(map[string]int)

	// Assertions
	assert.Equal(t, 0, len(emptyMap1))
	assert.Equal(t, 0, len(emptyMap2))
	assert.Equal(t, emptyMap1, emptyMap2)
}

// ## Map Literal
// A map literal allocates a Map and sets an initial
// number of key/value pairs in one statement.
func Test_Map_Literal(t *testing.T) {

	datesOfBirth := map[string]int{
		"Newton":   1643,
		"Faraday":  1791,
		"Einstein": 1879,
	}

	// Assertions
	assert.Equal(t, 1643, datesOfBirth["Newton"])
	assert.Equal(t, 1791, datesOfBirth["Faraday"])
	assert.Equal(t, 1879, datesOfBirth["Einstein"])
}

// ## Key/Value Insertion and Updating
// Key/value pairs are added (or updated) to a map using the
// `map[<KEY>] = <VALUE>` syntax.
func Test_Map_Insertion(t *testing.T) {
	datesOfBirth := map[string]int{}

	// Insertion of non-existing key/value pairs
	datesOfBirth["Newton"] = 1543
	datesOfBirth["Faraday"] = 1791
	datesOfBirth["Einstein"] = 1879

	// Assertions #1
	assert.Equal(t, 1543, datesOfBirth["Newton"])
	assert.Equal(t, 1791, datesOfBirth["Faraday"])
	assert.Equal(t, 1879, datesOfBirth["Einstein"])

	// Update of existing key/value pair
	datesOfBirth["Newton"] = 1643

	// Assertions #2
	assert.Equal(t, 1643, datesOfBirth["Newton"])
}

// ## Deletion
// Key/value pairs are deleted using the `delete(<MAP>, <KEY>)`
// function.
func Test_Map_Deletion(t *testing.T) {
	var datesOfBirth = map[string]int{
		"Newton":   1643,
		"Faraday":  1791,
		"Einstein": 1879,
	}

	delete(datesOfBirth, "Faraday")

	// Assertions
	assert.Equal(t, 0, datesOfBirth["Faraday"])
	assert.Equal(t, 2, len(datesOfBirth))

}

// ## Membership Test
// Key membership can be tested by checking the second return value when
// querying a key.
func Test_Map_Membership(t *testing.T) {

	// Set up Initial Map
	var datesOfBirth = map[string]int{
		"Newton":   1643,
		"Faraday":  1791,
		"Einstein": 1879,
	}

	// Test key membership
	_, ok1 := datesOfBirth["Newton"]
	_, ok2 := datesOfBirth["Curie"]

	// Assertions
	assert.Equal(t, true, ok1)
	assert.Equal(t, false, ok2)
}

// ## Getting All Keys and Values
// This is achived via iteration.
func Test_Map_Get_All(t *testing.T) {
	var datesOfBirth = map[string]int{
		"Newton":   1643,
		"Faraday":  1791,
		"Einstein": 1879,
	}

	// Get all keys
	var keys = ""
	for key := range datesOfBirth {
		keys = keys + key
	}

	// Get all values
	var values = ""
	for _, value := range datesOfBirth {
		values = values + strconv.Itoa(value)
	}

	// Assertions (Keys)
	assert.Equal(t, true, strings.Contains(keys, "Einstein"))
	assert.Equal(t, true, strings.Contains(keys, "Newton"))
	assert.Equal(t, true, strings.Contains(keys, "Faraday"))

	// Assertions (Values)
	assert.Equal(t, true, strings.Contains(values, "1879"))
	assert.Equal(t, true, strings.Contains(values, "1643"))
	assert.Equal(t, true, strings.Contains(values, "1791"))
}
