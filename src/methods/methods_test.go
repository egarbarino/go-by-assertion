// # Methods
// Methods are functions that operate against a given type and can, thus,
// be invoked using the dot notation against the relevant type using the
// form `<VAR>.<METHOD>()`.
//
// Ignore-On
package methods

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Method Declaration
// Methods are declared by defining a top level function which receives
// the target type and specifies a return type as a function type that
// represents the method in the form `func (<VAR> <TYPE>) <METHOD>`
// where `<TYPE>` is the struct or type alias to which we want to
// add a method, and `<METHOD>` is the full function signature for the method.
// In this case, `<VAR>` will be a value, so any changes to the type will be
// local to the function. If we want the method to affect the provided
// value, then the pointer declaration should be used by prefixing the
// type with a star: `*<TYPE>`.
//
// Let us now proceed with the example.
// First we declare a struct called `Employee` to which we will add
// two methods further on.
type Employee struct {
	Name   string
	Salary int
}

// Secondly, we add a method called `FailToIncreaseSalaryBy()` which allows
// raising the employee's salary by the provided amount and returns the
// new effective salary. It is prefixed with `Fail` since the method receives
// a value and the changes only occur on a copy, rather than on the original value.
func (e Employee) FailToIncreaseSalaryBy(extraSalary int) int {
	e.Salary = e.Salary + extraSalary // Only changed within this method's scope
	return e.Salary
}

// Thirdly, we add a method called `SucceedToIncreaseSalaryBy()` which uses
// the pointer notation and that, unlike the previous method, succeeds in
// modifying the originally provided value's attribute.
func (e *Employee) SucceedToIncreaseSalaryBy(extraSalary int) int {
	e.Salary = e.Salary + extraSalary // Passed value actually changed!
	return e.Salary
}

// Finally, we implement two sets of assertions to verify the differences
// between the two methods.
func Test_Method(t *testing.T) {
	e := Employee{"Jazmin", 2000}

	// Assertion #1
	// Value-wise Method Call
	var newSalary = e.FailToIncreaseSalaryBy(1000)
	assert.Equal(t, 3000, newSalary)
	assert.Equal(t, 2000, e.Salary) // Employee Salary not changed!

	// Assertion #2
	// Reference-wise Method Call (Side Effects!)
	e.SucceedToIncreaseSalaryBy(1000)
	assert.Equal(t, 3000, e.Salary) // Employee Salary changed!
}

// ## Struct Methods on Nested Structures
// Whenever a struct embeds other structs, a method may change
// the inner struct only, resulting in the modification of the outer struct
// as well.
//
// Let us see an example. First we declare the inner struct `Address`, and
// `PublicEmployee` which embeds it, and a method against `Address` to
// change the postcode called `changePostCode()`.
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

func (a *Address) ChangePostCode(newPostCode string) {
	a.Postcode = newPostCode // Passed value actually changed!
}

// We finally prove that when declaring a `PublicEmployee` value and changing
// the postcode of its embedded `Address` component, we effectively change the
// entire value.
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

	// Assertions
	assert.Equal(t, "SW18RLN", e.PrimaryAddress.Postcode) // Before change
	e.PrimaryAddress.ChangePostCode("TW18NLJ")            // change postcode!
	assert.Equal(t, "TW18NLJ", e.PrimaryAddress.Postcode) // After change

}

// ## Methods on Basic Types
// Methods cannot be declared directly against built-in types such as string.
// The solution is creating a _type alias_ using the `type <NAME> <TYPE>` notation.
type StringAlias string

func (s StringAlias) ToUpperCase() string { // Method against StringAlias
	return strings.ToUpper(string(s))
}
func Test_Method_On_Basic_Type(t *testing.T) {
	greeting := StringAlias("hello")

	// Assertions
	assert.Equal(t, "HELLO", greeting.ToUpperCase())
}
