package main

import (
	"fmt"
	"time"
)

func main() {
	// go sayHello("jon")
	// go sayHello("tom")
	// sayHello("arya stark")

	go sayWithTime(time.Second * 1)

	go sayWithTime(time.Second * 3)

	time.Sleep(time.Second * 5)
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayWithTime(d time.Duration) {
	time.Sleep(d)
	fmt.Println("Slept for", d)
}

func addOne(a int) int {
	return a + 1
}
