package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe          = false
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

// Constants are declared like variables, but with the const keyword
// Constants can be character, string, boolean, or numeric values
// Constants cannot be declared using the := syntax
const Pi = 3.14

// Numeric constants are high-precision values
// An untyped constant takes the type needed by its context
// Try printing needInt(Big) too
// (An int can store at maximum a 64-bit integer, and sometimes less)
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1 << 1, or 2
	Small = Big >> 99
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables declared without an explicit initial value are given their zero value
	// The zero value is:
	// 0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side
	// When the right hand side of the declaration is typed, the new variable is of that same type:
	// var i int
	// j := i // j is an int
	// But when the right hand side contains n untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
	// i := 42 // int
	// f := 3.142 // float64
	// g := 0.867 + 0.5i // complex128
	// Try changing code bellow
	v := 42 // Change me!
	fmt.Printf("v is of type %T\n", v)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return 1 + x*10
}
func needFloat(x float64) float64 {
	return x * 0.1
}

// The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements
// The int uint, and uintptr types are usually 32 bits wide on 32-bit system and 64 bits wide on 64-bit system. When you need an integer value you should use iny unless you have a specific reason to use a sized or unsigned integer type.
