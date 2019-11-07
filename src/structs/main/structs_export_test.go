// Then here, we have the results from the perspective
// of a file importing the package `model` presented above.

package main

import (
	"testing"

	"model"

	"github.com/stretchr/testify/assert"
)

func Test_Struct_Export_Example(t *testing.T) {
	//s = model.secretAgent{"Luigi", 1000} // Unexported
	//e = model.PublicEmployee{"Jazmin", "Doctor", 20000} Salary unexported!
	e := model.PublicEmployee{Name: "Jazmin", Role: "Doctor"}
	assert.Equal(t, "Jazmin", e.Name)
	assert.Equal(t, "Doctor", e.Role)
	// assert.Equal(t, 2000, e.Salary) Salary ununexported!
}
