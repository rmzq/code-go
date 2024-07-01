package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	u := e.Group("/users")
	u.GET("", findAllUsers)
	u.POST("", createUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func findAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"message": "find all user"})
}

func createUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"message": "create user"})
}
