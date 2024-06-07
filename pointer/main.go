package main

import "fmt"

func main() {
	var num int
	num = 5
	fmt.Println(num)
	num += 2
	num2 := &num
	fmt.Println(*num2)

	var testCoy int
	// testCoy = *num2
	fmt.Println(testCoy)
}
