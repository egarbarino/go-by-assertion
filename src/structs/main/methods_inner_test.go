// Ignore-On
package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Struct Methods on Nested Structures

type Address struct {
	Number     int
	StreetName string
	Postcode   string
}

type PublicEmployee struct {
	Name           string
	Salary         int
	PrimaryAddress Address
}

// Method gets a reference to Address (only)
func (a *Address) ChangePostCode(newPostCode string) {
	a.Postcode = newPostCode // Passed value actually changed!
}

func Test_Nested_Struct(t *testing.T) {

	e := PublicEmployee{
		Name:   "Jazmin",
		Salary: 2000,
		PrimaryAddress: Address{
			Number:     12,
			StreetName: "High street",
			Postcode:   "SW18RLN",
		},
	}

	assert.Equal(t, "SW18RLN", e.PrimaryAddress.Postcode)
	e.PrimaryAddress.ChangePostCode("TW18NLJ")
	assert.Equal(t, "TW18NLJ", e.PrimaryAddress.Postcode)

}

// ## Methods on Basic Types

// Create an alias first
type StringAlias string

// Create method using value reference (no asterisk)
func (s StringAlias) ToUpperCase() string {
	return strings.ToUpper(string(s))
}
func Test_Method_On_Basic_Type(t *testing.T) {
	greeting := StringAlias("hello")
	assert.Equal(t, "HELLO", greeting.ToUpperCase())
}
