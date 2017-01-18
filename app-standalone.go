package main

import (
	"github.com/kmkatsma/hello/products"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createMux() *echo.Echo {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Route => handler
	e.GET("/products/:id", products.Get)
	e.POST("/products", products.Post)

	return e
}

func main() {

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
