---
title: Go Examples By Unit Testing
author: Ernesto Garbarino
date: 2019-11-05
---
# Flow Control
## If-Then-Else


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L13)


``` go
var successCounter = 0
var failureCounter = 0
// Simple If
if 2*3 == 6 {
	successCounter++
}
// If Else
if 2*3 == 7 {
	failureCounter++
} else {
	successCounter++
}
// If Else If
if 2*3 == 8 {
	failureCounter++
} else if 2*3 == 7 {
	failureCounter++
} else if 2*3 == 6 {
	successCounter++
} else {
	failureCounter++
}
3 ⇔ successCounter
0 ⇔ failureCounter
```
## Switch (Normal)


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L41)


``` go
var result = ""
switch 1 + 1 {
case 0:
	result = "zero"
case 1:
	result = "one"
case 2:
	result = "two"
default: // Optional
	result = "some_number"
}
"two" ⇔ result
```
## Switch (Multi-Value Case)


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L57)


``` go
// Just cases
var result = ""
switch 2 * 3 {
case 0, 1, 2, 3, 4, 5, 6:
	result = "between_zero_and_six"
case 7:
	result = "seven"
default: // Optional
	result = "some_number"
}
"between_zero_and_six" ⇔ result
```
## Switch (Expression Cases)


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L72)


``` go
var result = ""
switch { // No expression here
case 2*3 == 1:
	result = "one"
case 2*3 == 6:
	result = "six"
default: // Optional
	result = "some_number"
}
"six" ⇔ result
```
## Switch (Type Cases)


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L86)


``` go
testType := func(i interface{}) string {
	switch t := i.(type) { // the t := assignment is optional
	case int:
		return "int"
	case string:
		return "string"
	default: // Optional
		return "other type: " + fmt.Sprintf("%T", t)
	}
}
"string" ⇔ testType("hello")
"int" ⇔ testType(45)
"other type: float64" ⇔ testType(1.53)
```
## Inifinite Loop (While True Loop)


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L103)


``` go
var counter = 0
for {
	counter = counter + 1
	if counter == 3 {
		break
	}
}
3 ⇔ counter
```
## While Loop


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L115)


``` go
var counter = 0
for counter < 3 {
	counter = counter + 1
}
3 ⇔ counter
```
## Iterate Over Array Elements


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L124)


``` go
var indexSum = 0
var sum = 0
for index, currentValue := range []int{1, 2, 3} {
	indexSum = indexSum + index // 0 + 1 + 2 = 3
	sum = sum + currentValue    // 1 + 2 + 3 = 6
}
3 ⇔ indexSum
6 ⇔ sum
```
## Iterate Over Array Elements But Use Only Index


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L136)


``` go
var indexSum = 0
for index := range []int{50, 999, 300} {
	indexSum = indexSum + index // 0 + 1 + 2 = 3
}
3 ⇔ indexSum
```
## Iterate Over Array Elements But Ignore Index


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L145)


``` go
var sum = 0
for _, currentValue := range []int{1, 2, 3} {
	sum = sum + currentValue // 1 + 2 + 3 = 6
}
6 ⇔ sum
```
## Iterate Over Keys and Values of a Map


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L154)


``` go
var keys, values string
for k, v := range map[string]string{"A": "(Argentina)", "B": "(Brazil)"} {
	keys = keys + k
	values = values + v
}
"AB" ⇔ keys
"(Argentina)(Brazil)" ⇔ values
```
## Iterate Over the Unicode Characters of a String


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L165)


``` go
var word string
for _, v := range "😊 olleh" {
	word = string(v) + word
}
"hello 😊" ⇔ word
```
## Iterate Over The Bytes of a String


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L174)


``` go
var reversedHello = "😊 olleh"
var word string
for i := 0; i < len(reversedHello); i++ {
	word = string(reversedHello[i]) + word
}
"hello \u008a\u0098\u009fð" ⇔ word
```
## Regular Foor Loop


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L184)


``` go
var counter = 0
for i := 0; i < 3; i++ {
	counter = counter + i // 0 + 1 + 2 = 3
}
3 ⇔ counter
```
## Defer


Source: [flowcontrol_test.go](https://github.com/egarbarino/go-examples/tree/master/src/flowcontrol/flowcontrol_test.go#L193)


``` go
var actions = ""
func addAction(action string) {
	actions = actions + action
}
func trackDefer() {
	addAction("1")
	defer addAction("[d1]")
	addAction("2")
	defer addAction("[d2]")
	addAction("3")
	defer addAction("[d3]")
	addAction("4")
	if true {
		return
	}
	defer addAction("[d4]")
	addAction("5")
}
trackDefer()
"1234[d3][d2][d1]" ⇔ actions
```
# Functions
## Functions With Multiple Return Values


Source: [functions_test.go](https://github.com/egarbarino/go-examples/tree/master/src/functions/functions_test.go#L12)


``` go
first, second := returnTwo()
'a' ⇔ first
false ⇔ second
onlyFirst, _ := returnTwo()
'a' ⇔ onlyFirst
_, onlySecond := returnTwo()
false ⇔ onlySecond
func returnTwo() (rune, bool) {
	return 'a', false
}
```
## Variadic Functions


Source: [functions_test.go](https://github.com/egarbarino/go-examples/tree/master/src/functions/functions_test.go#L30)


``` go
0 ⇔ sum()
1 ⇔ sum(1)
3 ⇔ sum(1, 2)
6 ⇔ sum(1, 2, 3)
func sum(numbers ...int) int {
	var result int = 0
	for _, n := range numbers {
		result = result + n
	}
	return result
}
```
## Anonymous Functions


Source: [functions_test.go](https://github.com/egarbarino/go-examples/tree/master/src/functions/functions_test.go#L47)


``` go
// Direct use
3 ⇔ operation("sum")(1, 2)
6 ⇔ operation("mul")(2, 3)
// Function extraction first
sum := operation("sum")
mul := operation("mul")
3 ⇔ sum(1, 2)
6 ⇔ mul(2, 3)
// Create Inline
3 ⇔ func(v1 int, v2 int) int { return v1 + v2 }(1, 2)
func operation(op string) func(int, int) int {
	switch op {
	case "sum":
		return func(v1 int, v2 int) int {
			return v1 + v2
		}
	case "mul":
		return func(v1 int, v2 int) int {
			return v1 * v2
		}
	default:
		panic("op should be either \"sum\" or \"mul\"")
	}
}
```
## Closures


Source: [functions_test.go](https://github.com/egarbarino/go-examples/tree/master/src/functions/functions_test.go#L79)


``` go
var counter = 0
incrementer := func() int {
	counter++
	return counter
}
1 ⇔ incrementer()
counter++
3 ⇔ incrementer()
incrementer()
5 ⇔ incrementer()
```
# Basic Types
## Boolean (Bool) and Logical Operators


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L15)


``` go
// Declaration ⇔
var MinBool bool = false
var MaxBool bool = true
// Logical Operators
false ⇔ !true
true ⇔ !false
true ⇔ true || false
false ⇔ MinBool || MaxBool
false ⇔ true && false
```
## Signed Integer 8


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L28)


``` go
// Declaration
var MinInt8 int8 = -128
var MaxInt8 int8 = 127
// Bounds
true ⇔ MinInt8 == math.MinInt8
true ⇔ MaxInt8 == math.MaxInt8
```
## Signed Integer 16


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L38)


``` go
// Declaration
var MinInt16 int16 = -32768
var MaxInt16 int16 = 32767
// Bounds
true ⇔ MinInt16 == math.MinInt16
true ⇔ MaxInt16 == math.MaxInt16
```
## Signed Integer 32 (Rune)


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L48)


``` go
// Declaration
var MinInt32 rune = -2147483648
var MinRune rune = MinInt32
var MaxInt32 int32 = 2147483647
var MaxRune int32 = MaxInt32
// Bounds
true ⇔ MinInt32 == math.MinInt32
true ⇔ MaxInt32 == math.MaxInt32
true ⇔ MinRune == math.MinInt32
true ⇔ MaxRune == math.MaxInt32
// Rune is an alias for int32
MaxRune ⇔ MinInt32-1
MinRune ⇔ MaxInt32+1
```
## Signed Integer 64


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L65)


``` go
// Declaration
var MinInt64 int64 = -9223372036854775808
var MaxInt64 int64 = 9223372036854775807
// Bounds
true ⇔ MinInt64 == math.MinInt64
true ⇔ MaxInt64 == math.MaxInt64
```
## Signed Integer (Implementation-Specific Size: 32 or 64)


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L75)


``` go
// Only on 64-bit architectures
var MinInt int = -9223372036854775808
var MaxInt int = 9223372036854775807
// Same as int64 on 64-bit architectures
MinInt ⇔ math.MinInt64
MaxInt ⇔ math.MaxInt64
// Overflow behaviour
MaxInt ⇔ MinInt-1
MinInt ⇔ MaxInt+1
```
## Unsigned Integer 8 (Byte)


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L88)


``` go
// Declaration
var MinUInt8 uint8 = 0
var MaxUInt8 uint8 = 255 // 2^8-1
var MinByte byte = MinUInt8
var MaxByte byte = MaxUInt8
// Bounds
true ⇔ MaxUInt8 == math.MaxUint8
// Overflow
true ⇔ math.MaxUint8 == MinUInt8-1
true ⇔ MinUInt8 == MaxUInt8+1
// byte is an alias for uint8
MaxByte ⇔ MinUInt8-1
MinByte ⇔ MaxUInt8+1
```
## Unsigned Integer 16


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L105)


``` go
// Decaration
var MinUInt16 uint16 = 0
var MaxUInt16 uint16 = 65535 // 2^16-1
// Bounds
true ⇔ MaxUInt16 == math.MaxUint16
// Overflow
MaxUInt16 ⇔ MinUInt16-1
MinUInt16 ⇔ MaxUInt16+1
```
## Unsigned Integer 32


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L117)


``` go
// Declaration
var MinUInt32 uint32 = 0
var MaxUInt32 uint32 = 4294967295 // 2^32-1
// Bounds
true ⇔ MaxUInt32 == math.MaxUint32
// Overflow
MaxUInt32 ⇔ MinUInt32-1
MinUInt32 ⇔ MaxUInt32+1
```
## Unsigned Integer 64


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L129)


``` go
// Declaration
var MinUInt64 uint64 = 0
var MaxUInt64 uint64 = 18446744073709551615 // 2^64-1
// Bounds
true ⇔ MaxUInt64 == math.MaxUint64
// Overflow
MaxUInt64 ⇔ MinUInt64-1
MinUInt64 ⇔ MaxUInt64+1
```
## Unsigned Integer (Implementation Specific Size: 32 or 64)


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L141)


``` go
// Declaration
var MinUint uint = 0
var MaxUint uint = 18446744073709551615
// Bounds: Same as uint64 on 64-bit architectures
true ⇔ math.MaxUint64 == MaxUint
// Overflow behaviour
MaxUint ⇔ MinUint-1
MinUint ⇔ MaxUint+1
```
## Unsigned Integer Pointer (Implementation Specific Size: 32 or 64)


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L153)


``` go
// Declaration
var MinUintptr uintptr = 0
var MaxUintptr uintptr = 18446744073709551615
// Bounds: Same as uint64 on 64-bit architectures
true ⇔ math.MaxUint64 == MaxUintptr
// Overflow behaviour
MaxUintptr ⇔ MinUintptr-1
MinUintptr ⇔ MaxUintptr+1
```
## Float 32


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L165)


``` go
// Declaration
var MinFloat32 float32 = -1.401298464324817e-45
var MaxFloat32 float32 = 3.4028234663852886e+38
// Bounds
true ⇔ MinFloat32 == -math.SmallestNonzeroFloat32
true ⇔ MaxFloat32 == math.MaxFloat32
```
## Float 64


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L175)


``` go
// Declaration
var MinFloat64 float64 = -5e-324
var MaxFloat64 float64 = 1.7976931348623157e+308
// Bounds
true ⇔ MinFloat64 == -math.SmallestNonzeroFloat64
true ⇔ MaxFloat64 == math.MaxFloat64
```
## Complex64


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L185)


``` go
// Declaration and Bounds
complex(1.401298464324817e-45, 1.401298464324817e-45) ⇔ complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32)
complex(3.4028234663852886e+38, 3.4028234663852886e+38) ⇔ complex(math.MaxFloat32, math.MaxFloat32)
```
## Complex 128


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L192)


``` go
// Declaration and Bounds
complex(5e-324, 5e-324) ⇔ complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64)
complex(1.7976931348623157e+308, 1.7976931348623157e+308) ⇔ complex(math.MaxFloat64, math.MaxFloat64)
```
## Number Literals


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L199)


``` go
255 ⇔ 0xFF              // Hex
255 ⇔ 0XFF              // Hex
255 ⇔ 0377              // Octal
255 ⇔ 0o377             // Octal
255 ⇔ 0O377             // Octal
255 ⇔ 0b11111111        // Binary
255 ⇔ 0B11111111        // Binary
0.5 ⇔ .5                // Float
0.5 ⇔ 00.5              // Float
0.0 ⇔ 0.                // Float
50.0 ⇔ 0.5e2            // Float w/Exponent
50.0 ⇔ 0.5E2            // Float w/Exponent
50.0 ⇔ 0.5E+2           // Float w/Exponent
0.005 ⇔ 0.5E-2          // Float w/Exponent
complex128(1+2i) ⇔ 1+2i // Complex Number
1000000 ⇔ 1_000_000     // Digit seperator
```
## Number Arithmetic Operators


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L219)


``` go
5 ⇔ 3+2                                 // Sum
1 ⇔ 3-2                                 // Subtraction
6 ⇔ 3*2                                 // Multiplication
1 ⇔ 3/2                                 // Integer Division
1.5 ⇔ 3.0/2.0                           // Float Division
2 ⇔ 5%3                                 // Modulo operator
2.0 ⇔ math.Sqrt(4)                      // Square Root
true ⇔ math.Pi > 3.14 && math.Pi < 3.15 // Pi Constant
// State mutating operators
var x = 3
x++ // Increase value by one
4 ⇔ x
x-- // Decrease value by one
3 ⇔ x
```
## Number Type Conversion


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L237)


``` go
// Upcasting (Value-Preserving)
var smallValue byte = 5
var smallNegValue int8 = -5
true ⇔ 9000+uint(smallValue) == 9005
true ⇔ 9000+int(smallNegValue) == 8995
true ⇔ 9000+uint16(smallValue) == 9005
true ⇔ 9000+int16(smallNegValue) == 8995
true ⇔ 900000000+uint32(smallValue) == 900000005
true ⇔ 900000000+int32(smallNegValue) == 899999995
true ⇔ 9000000000000000000+uint64(smallValue) == 9000000000000000005
true ⇔ 9000000000000000000+int64(smallNegValue) == 8999999999999999995
// Downcasting (Value-Destructive)
var bigValue uint64 = 18446744073709551615    // 2^64
var bigNegValue int64 = -9223372036854775808  // 2^63
var bigNegValue2 int64 = -4611686018427387904 // 2^62
var bigNegValue3 int64 = -4294967296          // 2^32
uint32(4294967295) ⇔ uint32(bigValue)
int32(-5) ⇔ int32(smallNegValue)
int32(0) ⇔ int32(bigNegValue)  // Truncated to zero
int32(0) ⇔ int32(bigNegValue2) // Truncated to zero
int32(0) ⇔ int32(bigNegValue3) // Truncated to zero
uint16(65535) ⇔ uint16(bigValue)
uint8(255) ⇔ uint8(bigValue)
// Float and Integers
var floatValue float32 = 1.5
float32(1.0) ⇔ float32(1)
1 ⇔ int(floatValue)
```
## Rune Literals


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L268)


``` go
'A' ⇔ rune(65)     // Rune value
'A' ⇔ '\101'       // Octal
'A' ⇔ '\x41'       // Hex
'A' ⇔ '\u0041'     // 16-bit Unicode
'A' ⇔ '\U00000041' // 32-bit Unicode
'😀' ⇔ '\U0001F600' // 32-bit Unicode
'\a' ⇔ '\x07'      // Bell
'\b' ⇔ '\x08'      // Backspace
'\f' ⇔ '\x0C'      // Form feed
'\n' ⇔ '\x0A'      // Line Feed
'\r' ⇔ '\x0D'      // Carriage Return
'\t' ⇔ '\x09'      // Horizontal Tab
'\v' ⇔ '\x0b'      // Vertial Tab
'\\' ⇔ '\x5c'      // Backslash
'\'' ⇔ '\x27'      // Single Quote
```
## String Conversion


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L287)


``` go
// The string() function converts a Unicode code point!
"A" ⇔ string(65)
"😀" ⇔ string(0X0001F600)
"AB" ⇔ string([]byte{65, 66})
// Integer to String
"15" ⇔ fmt.Sprintf("%d", 15)
"15" ⇔ strconv.Itoa(15)
"15" ⇔ strconv.FormatInt(15, 10) // Base 10
"f" ⇔ strconv.FormatInt(15, 16)  // Base 16
"     15" ⇔ fmt.Sprintf("%7d", 15)
"     15" ⇔ fmt.Sprintf("%7d", 15)
7 ⇔ len(fmt.Sprintf("%7d", 15))
// String to Integer
if n, err := strconv.Atoi("15"); err == nil {
	15 ⇔ n
} else {
	t.Error("Unable to convert")
}
if n, err := strconv.Atoi("ex1!%5"); err == nil {
	t.Error("Conversion was not supposed to be successful", n)
} else {
	err ⇎ nil
}
// Float to String
"1.567000" ⇔ fmt.Sprintf("%f", 1.567)
"1.57" ⇔ fmt.Sprintf("%.2f", 1.567)
"   1.57" ⇔ fmt.Sprintf("%7.2f", 1.567)
7 ⇔ len(fmt.Sprintf("%7.2f", 1.567))
// String to Float
if n, err := strconv.ParseFloat("1.567", 32); err == nil {
	1.5670000314712524 ⇔ n
} else {
	t.Error("Unable to convert")
}
if n, err := strconv.ParseFloat("ex1!%5.67", 32); err == nil {
	t.Error("Conersion was not supposed to be successful", n)
} else {
	err ⇎ nil
}
```
## Constants


Source: [basictypes_test.go](https://github.com/egarbarino/go-examples/tree/master/src/basictypes/basictypes_test.go#L331)


``` go
// Constants are immutable
const x byte = 5
const y = 7
// n++ // Won't compile
// Array-like like Slices are mutable and cannot be made constants
// const slice = []byte{'a', 'b', 'c'} // Won't compile
```
# Arrays
## Initialisation


Source: [arrays_test.go](https://github.com/egarbarino/go-examples/tree/master/src/arrays/arrays_test.go#L14)


``` go
var arrayEmpty1 [0]int
arrayEmpty2 := [...]int{}
arrayEmpty1 ⇔ arrayEmpty2
0 ⇔ len(arrayEmpty1)
```
## Array Literals


Source: [arrays_test.go](https://github.com/egarbarino/go-examples/tree/master/src/arrays/arrays_test.go#L25)


``` go
// Explicit length
arrayTwo := [2]string{"Einstein", "Newton"}
// No need to specify length
arrayThree := [...]string{"Einstein", "Newton", "Curie"}
2 ⇔ len(arrayTwo)
3 ⇔ len(arrayThree)
```
## Set Element's Values


Source: [arrays_test.go](https://github.com/egarbarino/go-examples/tree/master/src/arrays/arrays_test.go#L37)


``` go
var arrayOne [1]int
var arrayTwo [2]string
arrayTwo[0] = "Einstein"
arrayTwo[1] = "Newton"
1 ⇔ len(arrayOne)
"Einstein" ⇔ arrayTwo[0]
"Newton" ⇔ arrayTwo[1]
```
## Traversal


Source: [arrays_test.go](https://github.com/egarbarino/go-examples/tree/master/src/arrays/arrays_test.go#L51)


``` go
array := [3]byte{'a', 'b', 'c'}
var result = ""
var counter = 0
for i := 0; i < len(array); i++ {
	counter++
	result = result + string(array[i])
}
3 ⇔ counter
"abc" ⇔ result
```
## Pass By Value vs Pass By Reference


Source: [arrays_test.go](https://github.com/egarbarino/go-examples/tree/master/src/arrays/arrays_test.go#L65)


``` go
array := [3]byte{'a', 'b', 'c'}
// Copying an array produces a brand new one
array2 := array
array2[0] = 'x'
byte('a') ⇔ array[0] // Rather than x
// Functions receive an array copy, rather than a reference
changeFirstElement(array)
byte('a') ⇔ array[0] // Rather than y
func changeFirstElement(array [3]byte) {
	array[0] = 'y'
}
```
# Slices
## Initialisation


Source: [slices_test.go](https://github.com/egarbarino/go-examples/tree/master/src/slices/slices_test.go#L15)


``` go
var sliceEmpty1 []int
var sliceEmpty2 = make([]int, 0, 0)
var sliceEmpty3 = []int{}
sliceEmpty1 ⇔ sliceEmpty2
sliceEmpty2 ⇔ sliceEmpty3
0 ⇔ len(sliceEmpty1)
```
## Slice Literals


Source: [slices_test.go](https://github.com/egarbarino/go-examples/tree/master/src/slices/slices_test.go#L27)


``` go
slice := []string{"Einstein", "Newton"}
2 ⇔ len(slice)
```
## Set Element's Values


Source: [slices_test.go](https://github.com/egarbarino/go-examples/tree/master/src/slices/slices_test.go#L35)


``` go
slice := []byte{'a', 'b', 'c'}
byte('a') ⇔ slice[0]
slice[0] = 'z'
byte('z') ⇔ slice[0]
```
## Append


Source: [slices_test.go](https://github.com/egarbarino/go-examples/tree/master/src/slices/slices_test.go#L46)


``` go
var sliceInt []int
var sliceStr []string
var sliceIntPrime = append(sliceInt, 15)
var sliceStrPrime = append(sliceStr, "Einstein", "Newton")
1 ⇔ len(sliceIntPrime)
2 ⇔ len(sliceStrPrime)
15 ⇔ sliceIntPrime[0]
"Einstein" ⇔ sliceStrPrime[0]
"Newton" ⇔ sliceStrPrime[1]
```
## Append


Source: [slices_test.go](https://github.com/egarbarino/go-examples/tree/master/src/slices/slices_test.go#L63)


``` go
// Same length
src1 := []byte{'a', 'b', 'c'}
dst1 := []byte{'x', 'y', 'z'}
copy(dst1, src1)
[]byte{'a', 'b', 'c'} ⇔ dst1
dst1 ⇔ src1
// Source smaller than destination
src2 := []byte{'a', 'b'}
dst2 := []byte{'x', 'y', 'z'}
copy(dst2, src2)
[]byte{'a', 'b', 'z'} ⇔ dst2
dst2 ⇎ src2
// Source larger than destination
src3 := []byte{'a', 'b', 'c', 'd'}
dst3 := []byte{'x', 'y', 'z'}
copy(dst3, src3)
[]byte{'a', 'b', 'c'} ⇔ dst3
dst3 ⇎ src3
```
## Subsets (Slices, Splits, Trims)


Source: [slices_test.go](https://github.com/egarbarino/go-examples/tree/master/src/slices/slices_test.go#L87)


``` go
byte('e') ⇔ []byte("hello")[1]
[]byte("ello") ⇔ []byte("hello")[1:]
[]byte("hel") ⇔ []byte("hello")[:3]
[]byte("ell") ⇔ []byte("hello")[1:4]
[]byte("ell") ⇔ []byte("hello")[1:4]
[][]byte{[]byte("a"), []byte("b"), []byte("c")} ⇔ bytes.Split([]byte("a b c"), []byte(" "))
[]byte("hello") ⇔ bytes.Trim([]byte("😊hello🍺"), "😊🍺")
[]byte("hello") ⇔ bytes.Trim([]byte("😊hello🍺"), "😊🍺")
[]byte("hello") ⇔ bytes.TrimSpace([]byte("\t \t hello\n"))
```
## Traversal


Source: [slices_test.go](https://github.com/egarbarino/go-examples/tree/master/src/slices/slices_test.go#L101)


``` go
slice := []byte{'a', 'b', 'c'}
var result = ""
var counter = 0
for i := 0; i < len(slice); i++ {
	counter++
	result = result + string(slice[i])
}
3 ⇔ counter
"abc" ⇔ result
```
## Pass By Value vs Pass By Reference


Source: [slices_test.go](https://github.com/egarbarino/go-examples/tree/master/src/slices/slices_test.go#L115)


``` go
slice := []byte{'a', 'b', 'c'}
// A copy of a slice still points to the original slice
slice2 := slice
slice2[0] = 'x'
byte('x') ⇔ slice[0] // Rather than a
// Functions receive an slice pointer, rather than a copy
changeFirstElement(slice)
byte('y') ⇔ slice[0] // Rather than a or x
// A slice's slice still points to the original slice
lastTwoLetters := slice[1:]
lastTwoLetters[0] = 'z'
byte('z') ⇔ slice[1] // Rather than b
func changeFirstElement(slice []byte) {
	slice[0] = 'y'
}
```
# Strings
## Quantifiers (Length, Occurrances)


Source: [strings_test.go](https://github.com/egarbarino/go-examples/tree/master/src/strings/strings_test.go#L13)


``` go
0 ⇔ len("")
5 ⇔ len("hello")
1 ⇔ strings.Index("hello fellows", "ell")
7 ⇔ strings.LastIndex("hello fellows", "ell")
2 ⇔ strings.IndexByte("hello", 'l')
3 ⇔ strings.LastIndexByte("hello", 'l')
4 ⇔ strings.IndexRune("🍷😊🍺🍷", '😊')
2 ⇔ strings.IndexAny("hi😊!", "😊🍺🍷")
6 ⇔ strings.IndexAny("🤣hi😊!", "😊🍺🍷")
12 ⇔ strings.LastIndexAny("🍷😊🍺🍷", "🍷")
```
## Predicates (Substring, Starts/Ends With)


Source: [strings_test.go](https://github.com/egarbarino/go-examples/tree/master/src/strings/strings_test.go#L27)


``` go
true ⇔ strings.Contains("hello", "llo")
true ⇔ strings.ContainsAny("hello", "elxqwer")
true ⇔ strings.ContainsAny("hi😊!", "😊🍺🍷")
true ⇔ strings.HasPrefix("hello", "he")
true ⇔ strings.HasSuffix("hello", "lo")
```
## Subsets (Slices, Splits, Trims)


Source: [strings_test.go](https://github.com/egarbarino/go-examples/tree/master/src/strings/strings_test.go#L36)


``` go
"e" ⇔ string("hello"[1])
"ello" ⇔ "hello"[1:]
"hel" ⇔ "hello"[:3]
"ell" ⇔ "hello"[1:4]
"ell" ⇔ "hello"[1:4]
[]string{"a", "b", "c"} ⇔ strings.Split("a b c", " ")
"hello" ⇔ strings.Trim("😊hello🍺", "😊🍺")
"hello" ⇔ strings.Trim("😊hello🍺", "😊🍺")
"hello" ⇔ strings.TrimSpace("\t \t hello\n")
```
## Replace


Source: [strings_test.go](https://github.com/egarbarino/go-examples/tree/master/src/strings/strings_test.go#L49)


``` go
"hello 😊😊" ⇔ strings.Replace("hello :):)", ":)", "😊", -1)
"hello 😊:)" ⇔ strings.Replace("hello :):)", ":)", "😊", 1)
```
## Convert Case (Title, Upper, Lower)


Source: [strings_test.go](https://github.com/egarbarino/go-examples/tree/master/src/strings/strings_test.go#L55)


``` go
"This Is The End" ⇔ strings.Title("this is the end")
"HELLO" ⇔ strings.ToUpper("HellO")
"hello" ⇔ strings.ToLower("HellO")
```
## Join (Concat) and Repeat


Source: [strings_test.go](https://github.com/egarbarino/go-examples/tree/master/src/strings/strings_test.go#L62)


``` go
"Hello-world" ⇔ strings.Join([]string{"Hello", "world"}, "-")
"hihihi" ⇔ strings.Repeat("hi", 3)
```
## String Formatting


Source: [stringformatting_test.go](https://github.com/egarbarino/go-examples/tree/master/src/strings/stringformatting_test.go#L12)


``` go
"17" ⇔ fmt.Sprintf("%d", 17)             // Decimal
"11" ⇔ fmt.Sprintf("%x", 17)             // Hexadecimal
"21" ⇔ fmt.Sprintf("%o", 17)             // Octal
"10001" ⇔ fmt.Sprintf("%b", 17)          // Binary
"10001" ⇔ fmt.Sprintf("%b", 17)          // Binary
"1.750000" ⇔ fmt.Sprintf("%f", 1.75)     // Floating Point
"1.75" ⇔ fmt.Sprintf("%.2f", 1.75)       // Floating Point
"true" ⇔ fmt.Sprintf("%t", true)         // Boolean
"😊" ⇔ fmt.Sprintf("%c", '😊'))             // Rune (Unicode point
"hello" ⇔ fmt.Sprintf("%s", "hello"))     // Rune (Unicode point
"    o" ⇔ fmt.Sprintf("%5s", "o")        // Right Alignment
"h    " ⇔ fmt.Sprintf("%-5s", "h")       // Left Alignment
"'😊'" ⇔ fmt.Sprintf("%q", '😊')           // Quoted Rune
"'😊'" ⇔ fmt.Sprintf("%q", '😊')           // Quoted Rune
"\"hello\"" ⇔ fmt.Sprintf("%q", "hello") // Quoted String
"c64" ⇔ fmt.Sprintf("%v%v", "c", 64)     // Default String formatting
"int" ⇔ fmt.Sprintf("%T", 64)            // Inferred Value's Type
"%" ⇔ fmt.Sprintf("%%")                  // Literal Percent Sign
```
# Maps
## Empty Map


Source: [maps_test.go](https://github.com/egarbarino/go-examples/tree/master/src/maps/maps_test.go#L14)


``` go
var emptyMap1 = map[string]int{}
var emptyMap2 = make(map[string]int)
0 ⇔ len(emptyMap1)
0 ⇔ len(emptyMap2)
emptyMap1 ⇔ emptyMap2
```
## Create, Read, Update, Delete


Source: [maps_test.go](https://github.com/egarbarino/go-examples/tree/master/src/maps/maps_test.go#L24)


``` go
// Set up Initial Map
var datesOfBirth = map[string]int{
	"Newton":   1643,
	"Faraday":  1791,
	"Einstein": 1579, // Wrong
}
3 ⇔ len(datesOfBirth)
// Update Existing Key's Value
datesOfBirth["Einstein"] = 1879
3 ⇔ len(datesOfBirth)
1879 ⇔ datesOfBirth["Einstein"]
// Add a New Key/Value Pair
datesOfBirth["Darwin"] = 1809
4 ⇔ len(datesOfBirth)
1809 ⇔ datesOfBirth["Darwin"]
// Delete Key/Value Pair
delete(datesOfBirth, "Darwin")
3 ⇔ len(datesOfBirth)
// Not found keys return the underlying type's zero value
0 ⇔ datesOfBirth["Darwin"]
// Deletion is idempotent
delete(datesOfBirth, "Darwin")
```
## Memerbship Test, Get All Keys, Get All Values


Source: [maps_test.go](https://github.com/egarbarino/go-examples/tree/master/src/maps/maps_test.go#L55)


``` go
// Set up Initial Map
var datesOfBirth = map[string]int{
	"Newton":   1643,
	"Faraday":  1791,
	"Einstein": 1879,
}
// Test key membership
_, ok := datesOfBirth["Newton"]
true ⇔ ok
_, ok2 := datesOfBirth["Curie"]
false ⇔ ok2
// Get all keys
var keys = ""
for key := range datesOfBirth {
	keys = keys + key
}
true ⇔ strings.Contains(keys, "Einstein")
true ⇔ strings.Contains(keys, "Newton")
true ⇔ strings.Contains(keys, "Faraday")
// Get all values
var values = ""
for _, value := range datesOfBirth {
	values = values + strconv.Itoa(value)
}
true ⇔ strings.Contains(values, "1879")
true ⇔ strings.Contains(values, "1643")
true ⇔ strings.Contains(values, "1791")
```
# Pointers
## Referencing and Dereferencing Pointers


Source: [pointers_test.go](https://github.com/egarbarino/go-examples/tree/master/src/pointers/pointers_test.go#L13)


``` go
var householdIncome int = 1000
var householdMembers string = "Luigi"
// &var passes the variable's pointer rather than value
marryJazmin(&householdIncome, &householdMembers)
1000 ⇔ householdIncome
"Luigi & Jazmin" ⇔ householdIncome
func marryJazmin(income *int, members *string) {
	// *var takes a pointer reference's value
	*income = *income + 2000
	*members = *members + " and Jazmin"
}
```
## Allocate Memory For New Type and Obtain Pointer


Source: [pointers_test.go](https://github.com/egarbarino/go-examples/tree/master/src/pointers/pointers_test.go#L29)


``` go
var n = new(int)
0 ⇔ *n
*n++
1 ⇔ *n
// n is not a value!. We get a pointer address
"0x" ⇔ fmt.Sprintf("%v", n)[:2]
```
## Multile Pointers to The Same Value


Source: [pointers_test.go](https://github.com/egarbarino/go-examples/tree/master/src/pointers/pointers_test.go#L39)


``` go
var n int = 0
var n1, n2, n3 = &n, &n, &n
0 ⇔ n
0 ⇔ *n1
0 ⇔ *n2
0 ⇔ *n3
*n2++ // Any pointer to n (could have been n1 or n3)
1 ⇔ n
1 ⇔ *n1
1 ⇔ *n2
1 ⇔ *n3
```
# Structs
## Empty Struct Without Fields


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L13)


``` go
type Empty struct{}
"structs.Empty" ⇔ fmt.Sprintf("%T", Empty{})
```
## Struct Initialisation


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L19)


``` go
type Employee struct {
	Name   string
	Salary int
}
e1 := Employee{}
e2 := new(Employee)
e1 ⇔ *e2
"" ⇔ e1.Name
0 ⇔ e2.Salary
```
## Struct Literals


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L33)


``` go
type Employee struct {
	Name   string
	Salary int
}
// Positional Literal
e1 := Employee{"Jazmin", 2000}
"Jazmin" ⇔ e1.Name
2000 ⇔ e1.Salary
// Full, Field-wise Literal
e2 := Employee{
	Name:   "Jazmin",
	Salary: 2000,
}
"Jazmin" ⇔ e2.Name
2000 ⇔ e2.Salary
// Partial Literal
e3 := Employee{
	Name: "Jazmin",
	// Salary: 2000, Omitted!!
}
// e3 := Employee{"Jazmin"} Illegal!
"Jazmin" ⇔ e3.Name
0 ⇔ e3.Salary // Zero value!
// Partial Literal (II)
e4 := Employee{
	// Name: "Jazmin", Omitted!!
	Salary: 2000,
}
"" ⇔ e4.Name
2000 ⇔ e4.Salary // Zero value!
// In-line Literal
Employee{Name: "", Salary: 2000} ⇔ e4
```
## Struct Updates


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L75)


``` go
type Employee struct {
	Name   string
	Salary int
}
e := Employee{Name: "Jazmin", Salary: 2000}
e.Name = "Luigi"
e.Salary = 1000
"Luigi" ⇔ e.Name
1000 ⇔ e.Salary
```
## Anonymous Struct


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L91)


``` go
jazmin := struct {
	Name   string
	Salary int
}{
	Name:   "Jazmin",
	Salary: 2000,
}
"Jazmin" ⇔ jazmin.Name
2000 ⇔ jazmin.Salary
```
## Anonymous Fields


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L105)


``` go
type Employee struct {
	string
	int // types have to be distinct; can't repeat them
}
// Positional construction
e1 := Employee{"Jazmin", 2000}
// Field-wise construction (can use types as labels)
e2 := Employee{string: "Jazmin", int: 2000}
e1 ⇔ e2
```
## Pass By Reference vs Pass By Value


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L121)


``` go
type Employee struct {
	Name   string
	Salary int
}
// Obtain Value (Default)
v := Employee{Name: "Jazmin", Salary: 2000}
var v2, v3 = v, v
2000 ⇔ v.Salary
v.Salary++
2001 ⇔ v.Salary
v2.Salary++
2001 ⇔ v2.Salary
v3.Salary++
2001 ⇔ v3.Salary
2001 ⇔ v.Salary
// Obtain Reference Rather than Value
r := &Employee{Name: "Jazmin", Salary: 2000}
var r2, r3 = r, r
2000 ⇔ r.Salary
r.Salary++
2001 ⇔ r.Salary
r2.Salary++
2002 ⇔ r2.Salary
r3.Salary++
2003 ⇔ r3.Salary
2003 ⇔ r.Salary // r = r2 = r3
```
## Nested Fields


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L154)


``` go
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
"IG9IJA" ⇔ e.PrimaryAddress.Postcode
"TW200RA" ⇔ e.SecondaryAddress.Postcode
```
## Nested Fields (Same Name as Struct)


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L191)


``` go
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
"IG9IJA" ⇔ e.Address.Postcode
```
## Nested Fields (Inline)


Source: [structs_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_test.go#L219)


``` go
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
"IG9IJA" ⇔ e1.Address.Postcode
```
## Exported vs Unexported Structs
We first declare two struct types, `secretAgent`
and `PublicEmployee` under the `model` package.


Source: [structs_export.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/vendor/model/structs_export.go#L4)


``` go
package model
type secretAgent struct { // First letter lower case: Unexported!
	name   string
	salary int
}
type PublicEmployee struct { // First letter upper case: Exported!
	Name   string // exported
	Role   string // exported
	salary int    // unexported (lower case!)
}
```
Then here, we have the results from the perspective
of a file importing the package `model` presented above.


Source: [structs_export_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/structs_export_test.go#L3)


``` go
package main
import (
	"testing"
	"model"
	"github.com/stretchr/testify/assert"
)
//s = model.secretAgent{"Luigi", 1000} // Unexported
//e = model.PublicEmployee{"Jazmin", "Doctor", 20000} Salary unexported!
e := model.PublicEmployee{Name: "Jazmin", Role: "Doctor"}
"Jazmin" ⇔ e.Name
"Doctor" ⇔ e.Role
// 2000 ⇔ e.Salary Salary ununexported!
```
# Methods
## Struct Methods (Value-wise and Reference-wise)


Source: [methods_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/methods_test.go#L12)


``` go
type Employee struct {
	Name   string
	Salary int
}
```
Method gets a copy of Employee, not the actual reference


Source: [methods_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/methods_test.go#L18)


``` go
func (e Employee) FailToIncreaseSalaryBy(extraSalary int) int {
	e.Salary = e.Salary + extraSalary // Only changed within this method's scope
	return e.Salary
}
```
Method gets a reference to Employee


Source: [methods_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/methods_test.go#L24)


``` go
func (e *Employee) SucceedToIncreaseSalaryBy(extraSalary int) int {
	e.Salary = e.Salary + extraSalary // Passed value actually changed!
	return e.Salary
}
e := Employee{"Jazmin", 2000}
// Value-wise Method Call
var newSalary = e.FailToIncreaseSalaryBy(1000)
3000 ⇔ newSalary
2000 ⇔ e.Salary // Employee Salary not changed!
// Reference-wise Method Call (Side Effects!)
e.SucceedToIncreaseSalaryBy(1000)
3000 ⇔ e.Salary // Employee Salary changed!
```
## Struct Methods on Nested Structures


Source: [methods_inner_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/methods_inner_test.go#L12)


``` go
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
```
Method gets a reference to Address (only)


Source: [methods_inner_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/methods_inner_test.go#L25)


``` go
func (a *Address) ChangePostCode(newPostCode string) {
	a.Postcode = newPostCode // Passed value actually changed!
}
e := PublicEmployee{
	Name:   "Jazmin",
	Salary: 2000,
	PrimaryAddress: Address{
		Number:     12,
		StreetName: "High street",
		Postcode:   "SW18RLN",
	},
}
"SW18RLN" ⇔ e.PrimaryAddress.Postcode
e.PrimaryAddress.ChangePostCode("TW18NLJ")
"TW18NLJ" ⇔ e.PrimaryAddress.Postcode
```
## Methods on Basic Types
Create an alias first


Source: [methods_inner_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/methods_inner_test.go#L50)


``` go
type StringAlias string
```
Create method using value reference (no asterisk)


Source: [methods_inner_test.go](https://github.com/egarbarino/go-examples/tree/master/src/structs/main/methods_inner_test.go#L53)


``` go
func (s StringAlias) ToUpperCase() string {
	return strings.ToUpper(string(s))
}
greeting := StringAlias("hello")
"HELLO" ⇔ greeting.ToUpperCase()
```
# Files
## Create, Get Metadata, Rename, Copy, and Delete File


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L26)


``` go
// Create file (equivalent to "touch")
fileHandle, err := os.Create(tmpRoot + "example.txt")
if err != nil {
	t.Error("Can't create file: ", err)
}
fileHandle.Close()
// Get File Metadata
fileInfo, err2 := os.Stat(tmpRoot + "example.txt")
if err2 != nil {
	t.Error("Can't access file: ", err2)
}
"example.txt" ⇔ fileInfo.Name()
int64(0) ⇔ fileInfo.Size()
os.FileMode(0x1b6) ⇔ fileInfo.Mode()
true ⇔ fileInfo.ModTime().Year() >= 2019
false ⇔ fileInfo.IsDir()
// Rename File
err3 := os.Rename(tmpRoot+"example.txt", tmpRoot+"deleteme.txt")
if err3 != nil {
	t.Error("Can't delete file: ", err3)
}
// Delete File
err4 := os.Remove(tmpRoot + "deleteme.txt")
if err4 != nil {
	t.Error("Can't delete file: ", err4)
}
```
## OS: Write File


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L60)


``` go
// Open File for Writing
fileHandle, err := os.Create(tmpRoot + "hello.txt")
if err != nil {
	t.Error("Can't open file: ", err)
}
defer fileHandle.Close()
// Write File
bytes := []byte("Hello 😀")
count, err2 := fileHandle.Write(bytes) // WriteString() also available
if err2 != nil {
	t.Error("Can't write to file: ", err2)
}
fileHandle.Sync() // Flush changes
10 ⇔ count
```
## OS: Read File


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L79)


``` go
Test_Write_File(t)
// Open File
fileHandle, err := os.Open(tmpRoot + "hello.txt")
if err != nil {
	t.Error("Can't open file: ", err)
}
defer fileHandle.Close()
// Successful Read
bytes := make([]byte, 10)
count, err2 := fileHandle.Read(bytes)
if err2 != nil {
	t.Error("Can't open file: ", err2)
}
"Hello 😀" ⇔ string(bytes)
10 ⇔ count
// EOF Error (No more bytes left to read)
var oneByteArray = []byte{0}
count, err3 := fileHandle.Read(oneByteArray)
io.EOF ⇔ err3
```
## IOUtil: Write File (Quick)


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L105)


``` go
err := ioutil.WriteFile(tmpRoot+"hello.txt", []byte("Hello 😀"), 0666)
if err != nil {
	t.Error("Can't open file: ", err)
}
```
## IOUtil: Read File (Using File Handle)


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L114)


``` go
Test_Write_File_IOUtil_Quick(t)
// Open file for reading
fileHandle, err := os.Open(tmpRoot + "hello.txt")
if err != nil {
	log.Fatal(err)
}
// Read File
bytes, err2 := ioutil.ReadAll(fileHandle)
if err2 != nil {
	t.Error("Can't open file: ", err2)
}
"Hello 😀" ⇔ string(bytes)
10 ⇔ len(bytes)
```
## IOUtil: Read File (Quick)


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L135)


``` go
Test_Write_File_IOUtil_Quick(t)
bytes, err := ioutil.ReadFile(tmpRoot + "hello.txt")
if err != nil {
	t.Error("Can't open file: ", err)
}
"Hello 😀" ⇔ string(bytes)
10 ⇔ len(bytes)
```
## BufIO: Write File


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L149)


``` go
// Open File For Writing
fileHandle, err := os.Create(tmpRoot + "hello.txt")
if err != nil {
	t.Error("Can't open file: ", err)
}
defer fileHandle.Close()
// Write File
writer := bufio.NewWriter(fileHandle)
count, err2 := writer.WriteString("Hello 😀")
if err2 != nil {
	t.Error("Can't open file: ", err2)
}
writer.Flush()
10 ⇔ count
```
## BufIO: Read File


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L169)


``` go
// Open File For Reading
fileHandle, err := os.Open(tmpRoot + "hello.txt")
if err != nil {
	t.Error("Can't open file: ", err)
}
defer fileHandle.Close()
// Read File
reader := bufio.NewReader(fileHandle)
line, err2 := reader.ReadString('o')
if err2 != nil {
	t.Error("Can't open file: ", err2)
}
"Hello" ⇔ line
// EOF Error (No more lines left to read)
_, err3 := reader.ReadString('o')
io.EOF ⇔ err3
```
## BufIO: Scanner


Source: [files_test.go](https://github.com/egarbarino/go-examples/tree/master/src/files/files_test.go#L192)


``` go
// First generate example file
contents := []byte("line1\nline2\nline3\n")
err := ioutil.WriteFile(tmpRoot+"lines.txt", contents, 0666)
if err != nil {
	t.Error("Can't open file: ", err)
}
// Open File For Reading
fileHandle, err := os.Open(tmpRoot + "lines.txt")
if err != nil {
	t.Error("Can't open file: ", err)
}
defer fileHandle.Close()
// Use Scanner
input := bufio.NewScanner(fileHandle)
var lineBuffer = ""
for input.Scan() {
	lineBuffer = lineBuffer + input.Text()
}
"line1line2line3" ⇔ lineBuffer
```
# Input/Output Streams
## Interacting with Stdin, Stdout, and Stderr


Source: [iostreams.go](https://github.com/egarbarino/go-examples/tree/master/src/iostreams/iostreams.go#L14)


``` go
func main() {
	// Read all Stdin
	data, _ := ioutil.ReadAll(os.Stdin)
	// Convert data to upper case and write to Stdout
	count, _ := os.Stdout.WriteString(strings.ToUpper(string(data)))
	// Write the number of bytes read to Stderr
	os.Stderr.WriteString(fmt.Sprintf("Bytes read: %d", count))
}
```
## Interacting with Shell Commands


Source: [iostreams_test.go](https://github.com/egarbarino/go-examples/tree/master/src/iostreams/iostreams_test.go#L14)


``` go
// Specify command
cmd := exec.Command("go", "run", "iostreams.go")
// Specify stdin input
cmd.Stdin = strings.NewReader("hello")
// Declare a buffer to capture stdout from command
var stdout bytes.Buffer
cmd.Stdout = &stdout
// Declare a buffer to capture stderr from command
var stderr bytes.Buffer
cmd.Stderr = &stderr
// Run Command
err := cmd.Run()
if err != nil {
	t.Error("Error:", err)
}
// Verify stdout and stderr values
"HELLO" ⇔ stdout.String()
"Bytes read: 5" ⇔ stderr.String()
```
