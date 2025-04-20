package basics

import (
	"fmt"
	"math"
)

func main() {
	// variables declaration

	var a, b int = 10, 3

	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const p float64 = 22 / 7
	fmt.Println(p)

	const q float64 = 22 / 7.0
	fmt.Println(q)

	// overflow example

	var maxInt int64 = 9223372036854775807 // max value int64 type
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// overflow example unsigned integer

	var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// underflow example with floating point type

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
