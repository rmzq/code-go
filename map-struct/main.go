package main

import "fmt"

type myMapType map[string]interface{}

type Address struct {
	Country string
	City    string
}

func (m *myMapType) Test() *myMapType {
	return m
}

func (m *myMapType) get() string {
	return (*m)["name"].(string)
}

func main() {
	var person = map[string]string{
		"name": "Arya",
		"age":  "20",
	}

	fmt.Println(person["name"])

	myMap := map[string][]string{}

	myMap["country"] = []string{"indonesia", "singapore", "malaysia"}
	myMap["city"] = []string{"jakarta", "singapore", "kuala-lumpur"}

	fmt.Println(myMap)

	fmt.Println(myMap["country"][0])

	var m myMapType
	m = myMapType{
		"name": "Arya",
		"age":  "20",
		"address": Address{
			Country: "indonesia",
			City:    "jakarta",
		},
	}

	fmt.Println(m.Test())

}
