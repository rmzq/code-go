package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func main() {
	person := Person{
		ID:          1,
		Name:        "Jon Snow",
		PhoneNumber: "7598379",
	}

	jsonByte, _ := json.Marshal(person)
	fmt.Println(string(jsonByte))

	jsonString := `{"ID":2,"Name":"Emilia Clark","PhoneNumber":"8759379"}`
	var p Person
	var d any

	err := json.Unmarshal([]byte(jsonString), &p)

	if err != nil {
		panic(err)
	}

	fmt.Println(p)
	fmt.Println(d)
}
