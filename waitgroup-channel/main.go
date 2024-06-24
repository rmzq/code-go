package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		fmt.Println("Hello 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello 2")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello 3")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Done!")

}
