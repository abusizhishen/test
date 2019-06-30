package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware

	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type Info struct {
	Name string
	Age string
	Lover string
	Sex string
}

func post(c echo.Context) error {
	var info = Info{
		Name:c.Param("name"),
		Age:c.Param("age"),
		Sex:c.Param("sex"),
		Lover:c.Param("lover"),
	}

	return c.JSON(200,info)
}