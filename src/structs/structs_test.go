// # Structs
// Structs allow defining complex multi-attribute data structures made up of
// basic types as well as other structs too. A struct in Go is something
// in between a classic C struct and a JSON object in JavaScript (when
// instantiated).
//
// Ignore-On
package structs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Empty Struct Without Fields
// Structs may serve as a type container; there is no need to
// declare any attributes.
func Test_Structs(t *testing.T) {
	type Empty struct{}

	// Assertions
	assert.Equal(t, "structs.Empty", fmt.Sprintf("%T", Empty{}))
}

// ## Struct Initialisation
// Structs, once declared, may be initialised using the literal syntax, `Name{}`,
// or the `new()` function, in which case a pointer is returned rather than
// a value.
func Test_Basic_Structs(t *testing.T) {
	type Employee struct {
		Name   string
		Salary int
	}

	e1 := Employee{}
	e2 := new(Employee)

	// Assertions
	assert.Equal(t, e1, *e2)
	assert.Equal(t, "", e1.Name)
	assert.Equal(t, 0, e2.Salary)
}

// ## Struct Literals
// Struct literals allow initialising a struct and populating its attributes
// in one statement. Attributes may be defined by name, or in a positional basis.
func Test_Struct_Literals(t *testing.T) {

	type Employee struct {
		Name   string
		Salary int
	}

	// ------------------ Positional Literal ---------------------

	e1 := Employee{"Jazmin", 2000}

	// Assertions
	assert.Equal(t, "Jazmin", e1.Name)
	assert.Equal(t, 2000, e1.Salary)

	// ---------------- Attribute-wise Literal -------------------

	e2 := Employee{
		Name:   "Jazmin",
		Salary: 2000,
	}

	// Assertions
	assert.Equal(t, "Jazmin", e2.Name)
	assert.Equal(t, 2000, e2.Salary)

	// ------------------ Partial Literal (I) --------------------

	// Partial Literal
	e3 := Employee{
		Name: "Jazmin",
		// Salary: 2000, Omitted!!
	}
	// e3 := Employee{"Jazmin"} Illegal!

	// Assertions
	assert.Equal(t, "Jazmin", e3.Name)
	assert.Equal(t, 0, e3.Salary) // Zero value!

	// ----------------- Partial Literal (II) --------------------

	// Partial Literal (II)
	e4 := Employee{
		// Name: "Jazmin", Omitted!!
		Salary: 2000,
	}

	// Assertions
	assert.Equal(t, "", e4.Name) // Zero value!
	assert.Equal(t, 2000, e4.Salary)
	assert.Equal(t, Employee{Name: "", Salary: 2000}, e4) // In-line
}

// ## Struct Updates
// Structs are updated by referencing their attributes using a dot-based
// notation like in other C-like languages.
func Test_Struct_Updates(t *testing.T) {

	type Employee struct {
		Name   string
		Salary int
	}

	e := Employee{Name: "Jazmin", Salary: 2000}

	// Update attributes!
	e.Name = "Luigi"
	e.Salary = 1000

	// Assertions
	assert.Equal(t, "Luigi", e.Name)
	assert.Equal(t, 1000, e.Salary)

}

// ## Anonymous Struct
// An anonymous struct combines the type definition and its instantiation
// in one statement.
func Test_Struct_Anonymous(t *testing.T) {
	jazmin := struct {
		Name   string
		Salary int
	}{
		Name:   "Jazmin",
		Salary: 2000,
	}

	// Assertions
	assert.Equal(t, "Jazmin", jazmin.Name)
	assert.Equal(t, 2000, jazmin.Salary)
}

// ## Anonymous Fields
// Fields may be identified by their type (if different), avoiding
// the need to declare custom names.
func Test_Struct_Anonymous_Fields(t *testing.T) {
	type Employee struct {
		string
		int // types have to be distinct; can't repeat them
	}

	// Positional construction
	e1 := Employee{"Jazmin", 2000}
	// Field-wise construction (can use types as labels)
	e2 := Employee{string: "Jazmin", int: 2000}

	// Assertions
	assert.Equal(t, e1, e2)
}

// ## Pass By Reference vs Pass By Value
// Structs are passed by value by default.
func Test_Struct_Pass(t *testing.T) {
	type Employee struct {
		Name   string
		Salary int
	}

	// ------------- Obtain Value (Default) ---------------

	v := Employee{Name: "Jazmin", Salary: 2000}
	var v2, v3 = v, v

	// Assertions
	assert.Equal(t, 2000, v.Salary)
	v.Salary++
	assert.Equal(t, 2001, v.Salary)
	v2.Salary++
	assert.Equal(t, 2001, v2.Salary)
	v3.Salary++
	assert.Equal(t, 2001, v3.Salary)
	assert.Equal(t, 2001, v.Salary)

	// --------------- Obtain Reference  ------------------

	r := &Employee{Name: "Jazmin", Salary: 2000} // Reference!
	var r2, r3 = r, r
	assert.Equal(t, 2000, r.Salary)
	r.Salary++
	assert.Equal(t, 2001, r.Salary)
	r2.Salary++
	assert.Equal(t, 2002, r2.Salary)
	r3.Salary++
	assert.Equal(t, 2003, r3.Salary)
	assert.Equal(t, 2003, r.Salary) // r = r2 = r3
}

// ## Nested Fields
// A struct's attribute type may be another struct.
func Test_Struct_Nested_Fields(t *testing.T) {
	type Address struct {
		Number     int
		StreetName string
		Postcode   string
	}

	type Employee struct {
		Name             string
		Salary           int
		PrimaryAddress   Address
		SecondaryAddress Address
	}

	a2 := Address{
		Number:     109,
		StreetName: "Larksfield",
		Postcode:   "TW200RA",
	}

	e := Employee{
		Name:   "Jazmin",
		Salary: 2000,
		PrimaryAddress: Address{
			Number:     101,
			StreetName: "Buckhurst",
			Postcode:   "IG9IJA",
		},
		SecondaryAddress: a2,
	}

	// Assertions
	assert.Equal(t, "IG9IJA", e.PrimaryAddress.Postcode)
	assert.Equal(t, "TW200RA", e.SecondaryAddress.Postcode)
}

// ## Nested Fields (Same Name as Struct)
// A struct attribute that includes another struct type does
// not need to declare a custom name; the struct's type will be
// used as the default name.
func Test_Struct_Nested_Fields_By_Convention(t *testing.T) {

	type Address struct {
		Number     int
		StreetName string
		Postcode   string
	}
	type Employee struct {
		Address // This will become an attribute named Address
		Name    string
		Salary  int
	}

	e := Employee{
		Name:   "Jazmin",
		Salary: 2000,
		Address: Address{
			Number:     101,
			StreetName: "Buckhurst",
			Postcode:   "IG9IJA",
		},
	}

	// Assertions
	assert.Equal(t, "IG9IJA", e.Address.Postcode)

}

// ## Nested Fields (Inline)
// Nested struct types do not need to be declared externally,
// they may be inlined as well.
func Test_Struct_Nested_Fields_Inline(t *testing.T) {

	type Employee struct {
		Name    string
		Salary  int
		Address struct {
			Number     int
			StreetName string
			Postcode   string
		}
	}

	e1 := Employee{
		Name:   "Jazmin",
		Salary: 2000,
		Address: struct {
			Number     int
			StreetName string
			Postcode   string
		}{
			Number:     101,
			StreetName: "Buckhurst",
			Postcode:   "IG9IJA",
		},
	}

	// Assertions
	assert.Equal(t, "IG9IJA", e1.Address.Postcode)

}
