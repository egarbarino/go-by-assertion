// # Maps
// Ignore-On
package maps

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Empty Map
func Test_Maps(t *testing.T) {
	var emptyMap1 = map[string]int{}
	var emptyMap2 = make(map[string]int)
	assert.Equal(t, 0, len(emptyMap1))
	assert.Equal(t, 0, len(emptyMap2))
	assert.Equal(t, emptyMap1, emptyMap2)
}

// ## Create, Read, Update, Delete
func Test_Map_CRUD(t *testing.T) {

	// Set up Initial Map
	var datesOfBirth = map[string]int{
		"Newton":   1643,
		"Faraday":  1791,
		"Einstein": 1579, // Wrong
	}
	assert.Equal(t, 3, len(datesOfBirth))

	// Update Existing Key's Value
	datesOfBirth["Einstein"] = 1879
	assert.Equal(t, 3, len(datesOfBirth))
	assert.Equal(t, 1879, datesOfBirth["Einstein"])

	// Add a New Key/Value Pair
	datesOfBirth["Darwin"] = 1809
	assert.Equal(t, 4, len(datesOfBirth))
	assert.Equal(t, 1809, datesOfBirth["Darwin"])

	// Delete Key/Value Pair
	delete(datesOfBirth, "Darwin")
	assert.Equal(t, 3, len(datesOfBirth))
	// Not found keys return the underlying type's zero value
	assert.Equal(t, 0, datesOfBirth["Darwin"])
	// Deletion is idempotent
	delete(datesOfBirth, "Darwin")

}

// ## Memerbship Test, Get All Keys, Get All Values
func Test_Map_Views(t *testing.T) {

	// Set up Initial Map
	var datesOfBirth = map[string]int{
		"Newton":   1643,
		"Faraday":  1791,
		"Einstein": 1879,
	}

	// Test key membership
	_, ok := datesOfBirth["Newton"]
	assert.Equal(t, true, ok)
	_, ok2 := datesOfBirth["Curie"]
	assert.Equal(t, false, ok2)

	// Get all keys
	var keys = ""
	for key := range datesOfBirth {
		keys = keys + key
	}
	assert.Equal(t, true, strings.Contains(keys, "Einstein"))
	assert.Equal(t, true, strings.Contains(keys, "Newton"))
	assert.Equal(t, true, strings.Contains(keys, "Faraday"))

	// Get all values
	var values = ""
	for _, value := range datesOfBirth {
		values = values + strconv.Itoa(value)
	}
	assert.Equal(t, true, strings.Contains(values, "1879"))
	assert.Equal(t, true, strings.Contains(values, "1643"))
	assert.Equal(t, true, strings.Contains(values, "1791"))
}
