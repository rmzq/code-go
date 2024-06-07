package main

import "fmt"

var myName string

func main() {
	var a int
	a = 1
	// a = 3
	fmt.Println(a)
	// fmt.Printf("Type of a is %T\n", a)
	fmt.Println(add(1, 2))
	// println("Hello, World!")
	// println("Hello, World!")

	fmt.Println(myName)
	// print()
	// println()
	// fmt.Printf("Type of MyName is %T\n", MyName)
}

func add(a int, b int) int {
	return a + b
}
