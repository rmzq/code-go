package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	mySlice := []string{"a", "b", "c"}
	val, err := getValue(mySlice, 3)
	// fmt.Println(getValue(mySlice, 3))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(val)

}

func getValue(mySlice []string, index int) (string, error) {
	if index >= len(mySlice) {
		return "", errors.New("index out of bound")
	}
	return mySlice[index], nil
}
