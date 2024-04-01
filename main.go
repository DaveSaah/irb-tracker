package main

import (
	"github.com/boring-school-work/irb-tracker/handlers"
	"github.com/boring-school-work/irb-tracker/tmpl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("assets"))
	e.Renderer = tmpl.CreateTemplate()

	e.GET("/", handlers.IndexView)
	e.GET("/home", handlers.HomeView)

	register := e.Group("/register")
	register.GET("", handlers.RegisterView)
	register.POST("/email", handlers.RegisterEmailValidate)
	register.POST("/student", handlers.RegisterStudentView)
	register.POST("/passwd", handlers.RegisterPasswdView)
	register.POST("", handlers.RegisterUser)

	// start server with logger
	e.Logger.Fatal(e.Start(":42069"))
}
