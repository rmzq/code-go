package main

import "strconv"

func main() {
	var a any = 1
	b := add(a.(int), 2)
	println(b)

	number := "20"

	convertedNumber, err := strconv.Atoi(number)
	if err != nil {
		println("Error")
	}
	println(convertedNumber)
}

func add(a int, b int) int {
	return a + b
}
