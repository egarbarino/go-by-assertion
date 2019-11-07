// # Structs
// Ignore-On
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Empty Struct Without Fields
func Test_Structs(t *testing.T) {
	type Empty struct{}
	assert.Equal(t, "structs.Empty", fmt.Sprintf("%T", Empty{}))
}

// ## Struct Initialisation
func Test_Basic_Structs(t *testing.T) {
	type Employee struct {
		Name   string
		Salary int
	}
	e1 := Employee{}
	e2 := new(Employee)
	assert.Equal(t, e1, *e2)
	assert.Equal(t, "", e1.Name)
	assert.Equal(t, 0, e2.Salary)
}

// ## Struct Literals
func Test_Struct_Literals(t *testing.T) {

	type Employee struct {
		Name   string
		Salary int
	}

	// Positional Literal
	e1 := Employee{"Jazmin", 2000}
	assert.Equal(t, "Jazmin", e1.Name)
	assert.Equal(t, 2000, e1.Salary)

	// Full, Field-wise Literal
	e2 := Employee{
		Name:   "Jazmin",
		Salary: 2000,
	}
	assert.Equal(t, "Jazmin", e2.Name)
	assert.Equal(t, 2000, e2.Salary)

	// Partial Literal
	e3 := Employee{
		Name: "Jazmin",
		// Salary: 2000, Omitted!!
	}
	// e3 := Employee{"Jazmin"} Illegal!
	assert.Equal(t, "Jazmin", e3.Name)
	assert.Equal(t, 0, e3.Salary) // Zero value!

	// Partial Literal (II)
	e4 := Employee{
		// Name: "Jazmin", Omitted!!
		Salary: 2000,
	}
	assert.Equal(t, "", e4.Name)
	assert.Equal(t, 2000, e4.Salary) // Zero value!

	// In-line Literal
	assert.Equal(t, Employee{Name: "", Salary: 2000}, e4)
}

// ## Struct Updates
func Test_Struct_Updates(t *testing.T) {

	type Employee struct {
		Name   string
		Salary int
	}
	e := Employee{Name: "Jazmin", Salary: 2000}

	e.Name = "Luigi"
	e.Salary = 1000

	assert.Equal(t, "Luigi", e.Name)
	assert.Equal(t, 1000, e.Salary)

}

// ## Anonymous Struct
func Test_Struct_Anonymous(t *testing.T) {
	jazmin := struct {
		Name   string
		Salary int
	}{
		Name:   "Jazmin",
		Salary: 2000,
	}
	assert.Equal(t, "Jazmin", jazmin.Name)
	assert.Equal(t, 2000, jazmin.Salary)
}

// ## Anonymous Fields
func Test_Struct_Anonymous_Fields(t *testing.T) {

	type Employee struct {
		string
		int // types have to be distinct; can't repeat them
	}

	// Positional construction
	e1 := Employee{"Jazmin", 2000}
	// Field-wise construction (can use types as labels)
	e2 := Employee{string: "Jazmin", int: 2000}
	assert.Equal(t, e1, e2)

}

// ## Pass By Reference vs Pass By Value
func Test_Struct_Pass(t *testing.T) {

	type Employee struct {
		Name   string
		Salary int
	}

	// Obtain Value (Default)
	v := Employee{Name: "Jazmin", Salary: 2000}
	var v2, v3 = v, v
	assert.Equal(t, 2000, v.Salary)
	v.Salary++
	assert.Equal(t, 2001, v.Salary)
	v2.Salary++
	assert.Equal(t, 2001, v2.Salary)
	v3.Salary++
	assert.Equal(t, 2001, v3.Salary)
	assert.Equal(t, 2001, v.Salary)

	// Obtain Reference Rather than Value
	r := &Employee{Name: "Jazmin", Salary: 2000}
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

	assert.Equal(t, "IG9IJA", e.PrimaryAddress.Postcode)
	assert.Equal(t, "TW200RA", e.SecondaryAddress.Postcode)

}

// ## Nested Fields (Same Name as Struct)
func Test_Struct_Nested_Fields_By_Convention(t *testing.T) {

	type Address struct {
		Number     int
		StreetName string
		Postcode   string
	}
	type Employee struct {
		Address // This will become a field named Address
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

	assert.Equal(t, "IG9IJA", e.Address.Postcode)

}

// ## Nested Fields (Inline)
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

	assert.Equal(t, "IG9IJA", e1.Address.Postcode)

}
