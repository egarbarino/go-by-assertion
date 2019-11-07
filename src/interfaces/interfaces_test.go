// # Interfaces
// Ignore-On
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Define an Interface

// Step 1: Define a struct type (or primitive alias)

type employee struct {
	personName string
	salary     int
}

type contractor struct {
	limitedCompany string
	dailyRate      int
}

// Step 2: Define an interface as a collection of related methods

type costingInterface interface {
	getYearlyCost() int
	// getTaxCost
	// getNetCost
}

// Step 3: Implement the interface's methods

func (c contractor) getYearlyCost() int {
	return c.dailyRate * 5 * 40 // working weeks p/ year
}

func (e employee) getYearlyCost() int {
	return (e.salary * 103) / 100 // Pension contributions 3%
}

// Step 4: Use the interface as a variable type
func Test_Interface(t *testing.T) {
	// Use Interface as Variable Type
	var c1, c2 costingInterface
	c1 = contractor{"RipOff Ltd", 500}
	c2 = employee{"John Smith", 80000}

	// Assertions
	assert.Equal(t, 100000, c1.getYearlyCost())
	assert.Equal(t, 82400, c2.getYearlyCost())
}
