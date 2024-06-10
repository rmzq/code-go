package main

import "fmt"

func main() {

	sayHello("John Doe")
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)

	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(getSumAndDiff(10, 5))

	a, b := getSumAndDiff2(10, 5)
	fmt.Println(a, b)

	myAdd := getAdd(1, 2, 2)
	// myAddAgain := addAgain(2)
	fmt.Println(myAdd(addAgain))

}

// variadic function
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func addAgain(x int) int {
	return x * 2
}

// function dengan slice
func sum2(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// multiple argument and return
func getSumAndDiff(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// named return value
func getSumAndDiff2(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return
}

func getAdd(a int, b int, c int) func(callback func(int) int) int {
	return func(callback func(int) int) int {
		return a + b + callback(c)
	}
}
