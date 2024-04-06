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
	register.POST("/passwd", handlers.RegisterPasswdValidate)
	register.POST("", handlers.RegisterUser)

	e.GET("/login", handlers.LoginView)
	e.POST("/login", handlers.LoginUser)

	e.GET("/dashboard", handlers.DashboardView)
	e.GET("/projects", handlers.ProjectsView)
	e.GET("/activity", handlers.ActivityView)
	e.GET("/review", handlers.ReviewView)

	// start server with logger
	e.Logger.Fatal(e.Start(":42069"))
}
