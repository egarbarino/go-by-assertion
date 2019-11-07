// # Methods
// Ignore-On
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Struct Methods (Value-wise and Reference-wise)

type Employee struct {
	Name   string
	Salary int
}

// Method gets a copy of Employee, not the actual reference
func (e Employee) FailToIncreaseSalaryBy(extraSalary int) int {
	e.Salary = e.Salary + extraSalary // Only changed within this method's scope
	return e.Salary
}

// Method gets a reference to Employee
func (e *Employee) SucceedToIncreaseSalaryBy(extraSalary int) int {
	e.Salary = e.Salary + extraSalary // Passed value actually changed!
	return e.Salary
}
func Test_Method(t *testing.T) {
	e := Employee{"Jazmin", 2000}

	// Value-wise Method Call
	var newSalary = e.FailToIncreaseSalaryBy(1000)
	assert.Equal(t, 3000, newSalary)
	assert.Equal(t, 2000, e.Salary) // Employee Salary not changed!

	// Reference-wise Method Call (Side Effects!)
	e.SucceedToIncreaseSalaryBy(1000)
	assert.Equal(t, 3000, e.Salary) // Employee Salary changed!
}
