// # Interfaces
// Interfaces are a collection of method signatures that a type
// may implement.
//
// Ignore-On
package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Interface Definition
// An interface is defined using the below notation:
//
// ```
// type <NAME> interface {
//   <METHOD1>
//   <METHOD2>
//   ...
// }
// ```
//
// In the below example, we create two types; `employee` and `contractor`, then declare
// an interface called `costing`, and finally implement said interface for both of
// the types.

type employee struct {
	personName string
	salary     int
}

type contractor struct {
	limitedCompany string
	dailyRate      int
}

type costingInterface interface { // Define interface (IF)!
	getYearlyCost() int
	// ... other methods here.
}

func (e employee) getYearlyCost() int { // Implement IF for employee
	return (e.salary * 103) / 100 // Pension contributions 3%
}

func (c contractor) getYearlyCost() int { // Implement IF for contractor
	return c.dailyRate * 5 * 40 // working weeks p/ year
}

func Test_Interface(t *testing.T) {
	// Use interface as a regular variable Type
	var c1, c2 costingInterface
	c1 = contractor{"RipOff Ltd", 500}
	c2 = employee{"John Smith", 80000}

	// Assertions
	assert.Equal(t, 100000, c1.getYearlyCost())
	assert.Equal(t, 82400, c2.getYearlyCost())
}

// ## The Empty Interface
// The empty interface specifies zero methods and may hold
// values of any type.

func Test_Interface_2(t *testing.T) {
	// v can be any value type!
	tellMeYourType := func(v interface{}) string {
		switch v.(type) {
		case string:
			return "string"
		case int:
			return "int"
		default:
			return "other"
		}
	}

	// All value types are compatible with interface{}!
	var v1, v2, v3 interface{}
	v1 = "hello"
	v2 = 45
	v3 = 1.5

	// Assertions
	assert.Equal(t, "string", tellMeYourType(v1))
	assert.Equal(t, "int", tellMeYourType(v2))
	assert.Equal(t, "other", tellMeYourType(v3))
}
