package main

import (
	"context"
	"fmt"
)

type contextKey string

const usernameKey contextKey = "username"
const userIDKey contextKey = "userID"

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, usernameKey, "John")
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Println(ctx.Value(usernameKey))
}
