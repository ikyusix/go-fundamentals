package main

// Every go program is made up of packages
// This code groups the imports into a parenthesized, "factored" import statement

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problem.\n", math.Sqrt(7))
	// In go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.
	fmt.Println(math.Pi)
}
