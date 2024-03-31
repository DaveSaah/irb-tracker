package main

import (
	"github.com/boring-school-work/irb-tracker/tmpl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("assets"))
	e.Renderer = tmpl.CreateTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/home", func(c echo.Context) error {
		return c.Render(200, "home", nil)
	})

	// start server with logger
	e.Logger.Fatal(e.Start(":42069"))
}
