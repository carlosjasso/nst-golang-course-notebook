package main

import (
	"fmt"
	"math"
	common "nst-golang-course/00-common"
	models "nst-golang-course/02-basic-data-types/models"
	"strings"
)

func main() {
	common.LabelMessage("Bitwise operators")
	value := 0 << 1
	fmt.Printf("Value: %v\n", value)

	common.LabelMessage("Statically typed")
	var apples int32 = 2
	var bananas int64 = 4
	var fruits int = int(apples) + int(bananas)
	fmt.Printf("Fruits: %v\n", fruits)

	common.LabelMessage("floats")
	var enumber float64 = 6.6e-32
	fmt.Printf("enumber: %v\n", enumber)

	common.LabelMessage("Math section")
	//   Max Values
	fmt.Printf("Int64 max value: %v\n", math.MaxInt64)
	//   Constants
	fmt.Printf("Pi: %v\n", math.Pi)
	//   complex numbers
	var complex1 complex128 = complex(1, 2)
	var complex2 complex128 = complex(3, 4)
	fmt.Printf("Complex1: %v\n", complex1)
	fmt.Printf("Complex2: %v\n", complex2)
	fmt.Printf("Complex real %v * %v = %v\n", complex1, complex2, real(complex1*complex2))
	fmt.Printf("Complex img %v * %v = %v\n", complex1, complex2, imag(complex1*complex2))

	common.LabelMessage("booleans")
	var mybool bool = true // defaults to false
	fmt.Printf("MyBool: %v\n", mybool)
	mybool = !mybool
	fmt.Printf("!MyBool: %v\n", mybool)

	common.LabelMessage("strings")
	var myString string = "Hello"
	fmt.Printf("myString: %v\n", myString)
	fmt.Printf("myString last char ASCII code: %v\n", myString[4])
	fmt.Printf("myString last char: %v\n", string(myString[4]))
	myString = "Hello Mexico"
	fmt.Printf("myString sliced: %v\n", myString[0:5])
	fmt.Printf("myString sliced: %v\n", myString[6:])
	fmt.Printf("Mexico has suffix \"co\": %v\n", strings.HasSuffix(myString, "co"))

	// Constants. Go to definition above the main method
	common.LabelMessage("Constants as enums")
	fmt.Printf("Constant value: %v\n", models.Donas)
	fmt.Printf("Monday value: %v\n", models.Monday)
	fmt.Printf("Thursday value: %v\n", models.Thursday)
}
