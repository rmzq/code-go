package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// var m sync.Mutex

	var xChan = make(chan int, 1)

	wg.Add(10000)

	var x int // value of x is 0

	for i := 0; i < 10000; i++ {
		// wg.Add(1)
		go func() {
			defer wg.Done()
			// prevent race condition
			<-xChan
			// m.Lock()
			x++
			// m.Unlock()
		}()
	}

	wg.Wait()

	// println("result", x)
	// it should print 10000
	fmt.Println(fmt.Sprintf("value of x = %d", x))
}
