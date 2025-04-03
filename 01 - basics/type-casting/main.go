package main

import (
	"fmt"
	"strconv"
)

/*
Typecasting
- is the process of converting a value from one data type to another
- Go is strictly typed, so implicit type conversion does not happen automatically like in some other languages
- you must explicitly converty types using type conversion functions
Syntax
newType := TypeName(value)
- TypeName(value) will explictly convert value into the specified TypeName
*/

func main() {
	// Type Conversion Between Basic Types

	// Integer to Floating-Point Conversion
	var intNum int = 42
	var floatNum float64 = float64(intNum)

	fmt.Println("integer:", intNum)
	fmt.Println("convered to Float:", floatNum)

	// Floating-Point to Integer Conversion (Truncation)
	var floatNum2 float64 = 4.53
	var intNum2 int = int(floatNum2)

	fmt.Println("Float:", floatNum2)
	fmt.Println("convered to Integer:", intNum2)

	// Integer to String Conversion
	var intNum3 = 100
	str := strconv.Itoa(intNum3)
	fmt.Println("String:", str)

	// String to Integer Conversion
	var strX = "1234"
	num, err := strconv.Atoi(strX)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Integer", num)
	}

	// String to Float Conversion
	strY := "3.142"
	floatNum3, err := strconv.ParseFloat(strY, 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Float:", floatNum3)
	}

	// Integer to Boolean Conversion
	var num2 = 1
	var flag bool = bool(num2 != 0)
	fmt.Println(flag)

	// Type Conversion Between Composite Types

	// Byte Slice to String
	byteSlice := []byte{72, 101, 108, 108, 111}
	strZ := string(byteSlice)

	fmt.Println("String:", strZ)

	// String to Byte Slice
	strO := "GoLang"
	byteSlice2 := []byte(strO)
	fmt.Println("String:", byteSlice2)

}
