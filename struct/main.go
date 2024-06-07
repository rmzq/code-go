package main

import "fmt"

type University struct {
	Name  string
	City  string
	IsPTN bool
}

func main() {
	var university = University{
		Name:  "University of Indonesia",
		City:  "Depok",
		IsPTN: true,
	}

	fmt.Println(university)
	fmt.Println(university.Name)

}
