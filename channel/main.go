package main

import "fmt"

func main() {
	ch := make(chan int)
	msgCh := make(chan string)

	go func() {
		ch <- 1
		sendData(msgCh, "Hello")
	}()

	go func() {
		ch <- 2
		sendData(msgCh, "World")
	}()

	go func() {
		ch <- 3
		sendData(msgCh, "!!!")
	}()

	var data int
	data = <-ch
	fmt.Println(data)

	data = <-ch
	fmt.Println(data)

	data = <-ch
	fmt.Println(data)

	// deadlock
	// data = <-ch
	// fmt.Println(data)

	fmt.Println(receiveData(msgCh))

	fmt.Println(receiveData(msgCh))

	fmt.Println(receiveData(msgCh))

}

func sendData(ch chan<- string, data string) {
	ch <- data
}

func receiveData(ch <-chan string) string {
	data := <-ch
	return data
}
