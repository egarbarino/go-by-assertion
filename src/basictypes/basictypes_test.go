// # Basic Types
// Ignore-On
package basictypes

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ## Boolean (Bool) and Logical Operators
func Test_Type_Bool(t *testing.T) {
	// Declaration
	var MinBool bool = false
	var MaxBool bool = true

	// Logical Operators
	assert.Equal(t, false, !true)
	assert.Equal(t, true, !false)
	assert.Equal(t, true, true || false)
	assert.Equal(t, false, MinBool || MaxBool)
	assert.Equal(t, false, true && false)
}

// ## Signed Integer 8
func Test_Type_Int8(t *testing.T) {
	// Declaration
	var MinInt8 int8 = -128
	var MaxInt8 int8 = 127

	// Bounds
	assert.Equal(t, true, MinInt8 == math.MinInt8)
	assert.Equal(t, true, MaxInt8 == math.MaxInt8)
}

// ## Signed Integer 16
func Test_Type_Int16(t *testing.T) {

	// Declaration
	var MinInt16 int16 = -32768
	var MaxInt16 int16 = 32767

	// Bounds
	assert.Equal(t, true, MinInt16 == math.MinInt16)
	assert.Equal(t, true, MaxInt16 == math.MaxInt16)
}

// ## Signed Integer 32 (Rune)
func Test_Type_Int32(t *testing.T) {
	// Declaration
	var MinInt32 rune = -2147483648
	var MinRune rune = MinInt32
	var MaxInt32 int32 = 2147483647
	var MaxRune int32 = MaxInt32

	// Bounds
	assert.Equal(t, true, MinInt32 == math.MinInt32)
	assert.Equal(t, true, MaxInt32 == math.MaxInt32)
	assert.Equal(t, true, MinRune == math.MinInt32)
	assert.Equal(t, true, MaxRune == math.MaxInt32)

	// Rune is an alias for int32
	assert.Equal(t, MaxRune, MinInt32-1)
	assert.Equal(t, MinRune, MaxInt32+1)
}

// ## Signed Integer 64
func Test_Type_Int64(t *testing.T) {
	// Declaration
	var MinInt64 int64 = -9223372036854775808
	var MaxInt64 int64 = 9223372036854775807

	// Bounds
	assert.Equal(t, true, MinInt64 == math.MinInt64)
	assert.Equal(t, true, MaxInt64 == math.MaxInt64)
}

// ## Signed Integer (Implementation-Specific Size: 32 or 64)
func Test_Type_Int(t *testing.T) {
	// Only on 64-bit architectures
	var MinInt int = -9223372036854775808
	var MaxInt int = 9223372036854775807

	// Same as int64 on 64-bit architectures
	assert.Equal(t, MinInt, math.MinInt64)
	assert.Equal(t, MaxInt, math.MaxInt64)

	// Overflow behaviour
	assert.Equal(t, MaxInt, MinInt-1)
	assert.Equal(t, MinInt, MaxInt+1)
}

// ## Unsigned Integer 8 (Byte)
func Test_Type_UInt8(t *testing.T) {
	// Declaration
	var MinUInt8 uint8 = 0
	var MaxUInt8 uint8 = 255 // 2^8-1
	var MinByte byte = MinUInt8
	var MaxByte byte = MaxUInt8

	// Bounds
	assert.Equal(t, true, MaxUInt8 == math.MaxUint8)

	// Overflow
	assert.Equal(t, true, math.MaxUint8 == MinUInt8-1)
	assert.Equal(t, true, MinUInt8 == MaxUInt8+1)

	// byte is an alias for uint8
	assert.Equal(t, MaxByte, MinUInt8-1)
	assert.Equal(t, MinByte, MaxUInt8+1)
}

// ## Unsigned Integer 16
func Test_Type_UInt16(t *testing.T) {
	// Decaration
	var MinUInt16 uint16 = 0
	var MaxUInt16 uint16 = 65535 // 2^16-1

	// Bounds
	assert.Equal(t, true, MaxUInt16 == math.MaxUint16)

	// Overflow
	assert.Equal(t, MaxUInt16, MinUInt16-1)
	assert.Equal(t, MinUInt16, MaxUInt16+1)
}

// ## Unsigned Integer 32
func Test_Type_UInt32(t *testing.T) {
	// Declaration
	var MinUInt32 uint32 = 0
	var MaxUInt32 uint32 = 4294967295 // 2^32-1

	// Bounds
	assert.Equal(t, true, MaxUInt32 == math.MaxUint32)

	// Overflow
	assert.Equal(t, MaxUInt32, MinUInt32-1)
	assert.Equal(t, MinUInt32, MaxUInt32+1)
}

// ## Unsigned Integer 64
func Test_Type_UInt64(t *testing.T) {
	// Declaration
	var MinUInt64 uint64 = 0
	var MaxUInt64 uint64 = 18446744073709551615 // 2^64-1

	// Bounds
	assert.Equal(t, true, MaxUInt64 == math.MaxUint64)

	// Overflow
	assert.Equal(t, MaxUInt64, MinUInt64-1)
	assert.Equal(t, MinUInt64, MaxUInt64+1)
}

// ## Unsigned Integer (Implementation Specific Size: 32 or 64)
func Test_Type_UInteger(t *testing.T) {
	// Declaration
	var MinUint uint = 0
	var MaxUint uint = 18446744073709551615

	// Bounds: Same as uint64 on 64-bit architectures
	assert.Equal(t, true, math.MaxUint64 == MaxUint)

	// Overflow behaviour
	assert.Equal(t, MaxUint, MinUint-1)
	assert.Equal(t, MinUint, MaxUint+1)
}

// ## Unsigned Integer Pointer (Implementation Specific Size: 32 or 64)
func Test_Type_UIntegerPointer(t *testing.T) {
	// Declaration
	var MinUintptr uintptr = 0
	var MaxUintptr uintptr = 18446744073709551615

	// Bounds: Same as uint64 on 64-bit architectures
	assert.Equal(t, true, math.MaxUint64 == MaxUintptr)

	// Overflow behaviour
	assert.Equal(t, MaxUintptr, MinUintptr-1)
	assert.Equal(t, MinUintptr, MaxUintptr+1)
}

// ## Float 32
func Test_Type_Float32(t *testing.T) {
	// Declaration
	var MinFloat32 float32 = -1.401298464324817e-45
	var MaxFloat32 float32 = 3.4028234663852886e+38

	// Bounds
	assert.Equal(t, true, MinFloat32 == -math.SmallestNonzeroFloat32)
	assert.Equal(t, true, MaxFloat32 == math.MaxFloat32)
}

// ## Float 64
func Test_Type_Float64(t *testing.T) {
	// Declaration
	var MinFloat64 float64 = -5e-324
	var MaxFloat64 float64 = 1.7976931348623157e+308

	// Bounds
	assert.Equal(t, true, MinFloat64 == -math.SmallestNonzeroFloat64)
	assert.Equal(t, true, MaxFloat64 == math.MaxFloat64)
}

// ## Complex64
func Test_Type_Complex64(t *testing.T) {
	// Declaration and Bounds
	assert.Equal(t, complex(1.401298464324817e-45, 1.401298464324817e-45), complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32))
	assert.Equal(t, complex(3.4028234663852886e+38, 3.4028234663852886e+38), complex(math.MaxFloat32, math.MaxFloat32))
}

// ## Complex 128
func Test_Type_Complex128(t *testing.T) {
	// Declaration and Bounds
	assert.Equal(t, complex(5e-324, 5e-324), complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64))
	assert.Equal(t, complex(1.7976931348623157e+308, 1.7976931348623157e+308), complex(math.MaxFloat64, math.MaxFloat64))
}

// ## Number Literals
func Test_Number_Literal(t *testing.T) {
	// Assertions
	assert.Equal(t, 255, 0xFF)              // Hex
	assert.Equal(t, 255, 0XFF)              // Hex
	assert.Equal(t, 255, 0377)              // Octal
	assert.Equal(t, 255, 0o377)             // Octal
	assert.Equal(t, 255, 0O377)             // Octal
	assert.Equal(t, 255, 0b11111111)        // Binary
	assert.Equal(t, 255, 0B11111111)        // Binary
	assert.Equal(t, 0.5, .5)                // Float
	assert.Equal(t, 0.5, 00.5)              // Float
	assert.Equal(t, 0.0, 0.)                // Float
	assert.Equal(t, 50.0, 0.5e2)            // Float w/Exponent
	assert.Equal(t, 50.0, 0.5E2)            // Float w/Exponent
	assert.Equal(t, 50.0, 0.5E+2)           // Float w/Exponent
	assert.Equal(t, 0.005, 0.5E-2)          // Float w/Exponent
	assert.Equal(t, complex128(1+2i), 1+2i) // Complex Number
	assert.Equal(t, 1000000, 1_000_000)     // Digit seperator
}

// ## Number Arithmetic Operators
func Test_Number_Operators(t *testing.T) {
	// Assertions
	assert.Equal(t, 5, 3+2)                                 // Sum
	assert.Equal(t, 1, 3-2)                                 // Subtraction
	assert.Equal(t, 6, 3*2)                                 // Multiplication
	assert.Equal(t, 1, 3/2)                                 // Integer Division
	assert.Equal(t, 1.5, 3.0/2.0)                           // Float Division
	assert.Equal(t, 2, 5%3)                                 // Modulo operator
	assert.Equal(t, 2.0, math.Sqrt(4))                      // Square Root
	assert.Equal(t, true, math.Pi > 3.14 && math.Pi < 3.15) // Pi Constant

	// State mutating operators
	var x = 3
	x++ // Increase value by one
	assert.Equal(t, 4, x)
	x-- // Decrease value by one
	assert.Equal(t, 3, x)
}

// ## Number Type Conversion
func Test_Number_Type_Conversion(t *testing.T) {
	// Upcasting (Value-Preserving)
	var smallValue byte = 5
	var smallNegValue int8 = -5
	assert.Equal(t, true, 9000+uint(smallValue) == 9005)
	assert.Equal(t, true, 9000+int(smallNegValue) == 8995)
	assert.Equal(t, true, 9000+uint16(smallValue) == 9005)
	assert.Equal(t, true, 9000+int16(smallNegValue) == 8995)
	assert.Equal(t, true, 900000000+uint32(smallValue) == 900000005)
	assert.Equal(t, true, 900000000+int32(smallNegValue) == 899999995)
	assert.Equal(t, true, 9000000000000000000+uint64(smallValue) == 9000000000000000005)
	assert.Equal(t, true, 9000000000000000000+int64(smallNegValue) == 8999999999999999995)

	// Downcasting (Value-Destructive)
	var bigValue uint64 = 18446744073709551615    // 2^64
	var bigNegValue int64 = -9223372036854775808  // 2^63
	var bigNegValue2 int64 = -4611686018427387904 // 2^62
	var bigNegValue3 int64 = -4294967296          // 2^32
	assert.Equal(t, uint32(4294967295), uint32(bigValue))
	assert.Equal(t, int32(-5), int32(smallNegValue))
	assert.Equal(t, int32(0), int32(bigNegValue))  // Truncated to zero
	assert.Equal(t, int32(0), int32(bigNegValue2)) // Truncated to zero
	assert.Equal(t, int32(0), int32(bigNegValue3)) // Truncated to zero
	assert.Equal(t, uint16(65535), uint16(bigValue))
	assert.Equal(t, uint8(255), uint8(bigValue))

	// Float and Integers
	var floatValue float32 = 1.5
	assert.Equal(t, float32(1.0), float32(1))
	assert.Equal(t, 1, int(floatValue))
}

// ## Rune Literals
func Test_Rune_Literals(t *testing.T) {
	// Assertions
	assert.Equal(t, 'A', rune(65))     // Rune value
	assert.Equal(t, 'A', '\101')       // Octal
	assert.Equal(t, 'A', '\x41')       // Hex
	assert.Equal(t, 'A', '\u0041')     // 16-bit Unicode
	assert.Equal(t, 'A', '\U00000041') // 32-bit Unicode
	assert.Equal(t, '😀', '\U0001F600') // 32-bit Unicode
	assert.Equal(t, '\a', '\x07')      // Bell
	assert.Equal(t, '\b', '\x08')      // Backspace
	assert.Equal(t, '\f', '\x0C')      // Form feed
	assert.Equal(t, '\n', '\x0A')      // Line Feed
	assert.Equal(t, '\r', '\x0D')      // Carriage Return
	assert.Equal(t, '\t', '\x09')      // Horizontal Tab
	assert.Equal(t, '\v', '\x0b')      // Vertial Tab
	assert.Equal(t, '\\', '\x5c')      // Backslash
	assert.Equal(t, '\'', '\x27')      // Single Quote
}

// ## String Conversion
func Test_String_Conversion(t *testing.T) {
	// The string() function converts a Unicode code point!
	assert.Equal(t, "A", string(65))
	assert.Equal(t, "😀", string(0X0001F600))
	assert.Equal(t, "AB", string([]byte{65, 66}))

	// Integer to String
	assert.Equal(t, "15", fmt.Sprintf("%d", 15))
	assert.Equal(t, "15", strconv.Itoa(15))
	assert.Equal(t, "15", strconv.FormatInt(15, 10))    // Base 10
	assert.Equal(t, "f", strconv.FormatInt(15, 16))     // Base 16
	assert.Equal(t, "     15", fmt.Sprintf("%7d", 15))  // Padding
	assert.Equal(t, "15     ", fmt.Sprintf("%-7d", 15)) // Padding
	assert.Equal(t, 7, len(fmt.Sprintf("%7d", 15)))

	// String to Integer
	if n, err := strconv.Atoi("15"); err == nil {
		assert.Equal(t, 15, n)
	} else {
		t.Error("Unable to convert")
	}
	if n, err := strconv.Atoi("ex1!%5"); err == nil {
		t.Error("Conversion was not supposed to be successful", n)
	} else {
		assert.NotEqual(t, err, nil)
	}

	// Float to String
	assert.Equal(t, "1.567000", fmt.Sprintf("%f", 1.567))
	assert.Equal(t, "1.57", fmt.Sprintf("%.2f", 1.567))
	assert.Equal(t, "   1.57", fmt.Sprintf("%7.2f", 1.567))
	assert.Equal(t, 7, len(fmt.Sprintf("%7.2f", 1.567)))

	// String to Float
	if n, err := strconv.ParseFloat("1.567", 32); err == nil {
		assert.Equal(t, 1.5670000314712524, n)
	} else {
		t.Error("Unable to convert")
	}
	if n, err := strconv.ParseFloat("ex1!%5.67", 32); err == nil {
		t.Error("Conersion was not supposed to be successful", n)
	} else {
		assert.NotEqual(t, err, nil)
	}
}

// ## Constants
func Test_Constants(t *testing.T) {

	// Constants are immutable
	const x byte = 5
	const y = 7

	// x++ // Won't compile

	// Array-like like Slices are mutable and cannot be made constants
	// const slice = []byte{'a', 'b', 'c'} // Won't compile

}
