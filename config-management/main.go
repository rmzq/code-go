package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	fmt.Println(os.Getenv("INI_KEY"))
	println("Hello, World!")
}
