package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// main function
func main() {
	// create a new echo instance
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}