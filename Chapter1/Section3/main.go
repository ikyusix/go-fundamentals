package main

import "fmt"

// The var statement declares a list of variables; as in function argument lists, the type is last
// A var statement can be at package or function level. We see both in this example
var c, python, java bool

// A var declaration can include initializers, one per variable
// if an initializer is present, the type can be omitted; the variable will take the type of the initializer
// Variables with initializers sample code
var j, k int = 1, 2

func main() {
	var i int
	var m, n int = 1, 2
	o := 3
	p, PHP, MySQL := true, false, "no!"

	fmt.Println(i, c, python, java)

	// Variables with initializers sample code
	var l, ruby, kotlin = true, false, "no!"
	fmt.Println(j, k, l, ruby, kotlin)

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
	// Outside a function every statement begins with a keyword (var, func, and so om) and so the := construct is not available
	fmt.Println(m, n, o, p, PHP, MySQL)
}
