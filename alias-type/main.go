package main

import (
	"fmt"
	"strings"
)

type PersonName string

func main() {
	var name PersonName
	name = "John Doe"

	fmt.Println(name.GetFirstName(), name.GetLastName())
}

// untuk mendapatkan nama depan
func (p *PersonName) GetFirstName() string {
	return strings.Split(string(*p), " ")[0]
}

// untuk mendapatkan nama belakang
func (p *PersonName) GetLastName() string {
	return strings.Split(string(*p), " ")[1]
}
