package main

import "fmt"

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(addNd(42, 13))

	// Swap sample code
	a, b := swap("Hello", "world")
	fmt.Println(a, b)

	// Return values sample code
	fmt.Println(split(17))
}

// A function can tale zero or more arguments.
// In this example, add takes two parameters of type int.
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last
func add(x int, y int) int {
	return x + y
}
// In this example, we shortened
func addNd(x, y int) int {
	return x + y
}

// A function can return any number of result.
// The swap function returns two strings
func swap(x, y string) (string, string) {
	return y, x
}

// A return statement without arguments returns the named return values. This is known as a "naked" return
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
