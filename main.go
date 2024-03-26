package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func readyToLearn(c echo.Context) error {
	return c.String(http.StatusOK, "Ready to learn!")
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Define routes
	e.GET("/", readyToLearn)

	// Start the server
	e.Start(":8000")
}
