package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	midl := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			log.Println("before handler")

			err := next(c)

			log.Println("after handler 1")

			return err
		}
	}

	e.GET("/", func(c echo.Context) error {
		log.Println("inside handler")
		return c.String(200, "Hello, World! eh d")
	}, midl)

	e.Start(":1323")
}
