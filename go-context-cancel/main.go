package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	go doSomething(ctx)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("process has been 5 seconds, timeout!!!")
		cancel()
	})

	var s string
	fmt.Scan(&s)
}

func doSomething(ctx context.Context) {
	for {
		if isContextHasBeenCanceled(ctx) {
			fmt.Println("process has been timeout!")
			break
		}

		fmt.Println("do some work...")
		time.Sleep(time.Second)
	}
}

func isContextHasBeenCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
