package main

import (
	"fmt"
)

func main() {
	// a := 1
	name := "Arya"
	age := 20

	fmt.Printf("Hello, %s. You are %d years old.\n", name, age)
	// log.Fatal("Error")
	// fmt.Println(a)
	// fmt.Println(fmt.Sprintf("Heelo " + name + " " + age))

	// var
	// fmt.Scan()

	// var kalimat = "Hello, World! " + name + " " + age

	// fmt.Println(kalimat)

	var hobby string

	fmt.Scan(&hobby)

	fmt.Println(hobby)

}
