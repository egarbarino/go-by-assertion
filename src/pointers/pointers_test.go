// # Pointers
// Ignore-On
package pointers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Referencing and Dereferencing Pointers
func Test_Pointer_References(t *testing.T) {
	var householdIncome int = 1000
	var householdMembers string = "Luigi"
	// &var passes the variable's pointer rather than value
	marryJazmin(&householdIncome, &householdMembers)
	assert.Equal(t, 1000, householdIncome)
	assert.Equal(t, "Luigi & Jazmin", householdIncome)
}

func marryJazmin(income *int, members *string) {
	// *var takes a pointer reference's value
	*income = *income + 2000
	*members = *members + " and Jazmin"
}

// ## Allocate Memory For New Type and Obtain Pointer
func Test_Pointer_New(t *testing.T) {
	var n = new(int)
	assert.Equal(t, 0, *n)
	*n++
	assert.Equal(t, 1, *n)
	// n is not a value!. We get a pointer address
	assert.Equal(t, "0x", fmt.Sprintf("%v", n)[:2])
}

// ## Multile Pointers to The Same Value
func Test_Pointer_Bheaviour(t *testing.T) {
	var n int = 0
	var n1, n2, n3 = &n, &n, &n
	assert.Equal(t, 0, n)
	assert.Equal(t, 0, *n1)
	assert.Equal(t, 0, *n2)
	assert.Equal(t, 0, *n3)
	*n2++ // Any pointer to n (could have been n1 or n3)
	assert.Equal(t, 1, n)
	assert.Equal(t, 1, *n1)
	assert.Equal(t, 1, *n2)
	assert.Equal(t, 1, *n3)

}
