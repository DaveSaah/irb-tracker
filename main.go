package main

import (
	"github.com/boring-school-work/irb-tracker/handlers"
	"github.com/boring-school-work/irb-tracker/tmpl"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Static("assets"))
	e.Renderer = tmpl.CreateTemplate()

	e.Logger.SetLevel(log.DEBUG)

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

	e.POST("/logout", handlers.LogoutUser)

	search := e.Group("/search")
	search.POST("/projects", handlers.ProjectSearch)
	search.POST("/review", handlers.ReviewSearch)

	e.GET("/proposal", handlers.ProposalView)
	e.POST("/proposal", handlers.SubmitProposal)

	project := e.Group("/project")
	project.GET("", handlers.ProjectView)
	project.GET("/info", handlers.ProjectInfo)
	project.GET("/messages", handlers.ProjectMsg)
	project.GET("/brief", handlers.ProjectBrief)
	project.GET("/timeline", handlers.ProjectTimeline)
	project.GET("/participants", handlers.ProjectParticipants)
	project.GET("/review", handlers.ProjectReview)

	review := e.Group("/review")
	review.GET("", handlers.ReviewView)
	review.GET("/info", handlers.ReviewInfo)
	review.GET("/messages", handlers.ReviewMsg)
	review.GET("/brief", handlers.ReviewBrief)
	review.GET("/timeline", handlers.ReviewTimeline)
	review.GET("/participants", handlers.ReviewParticipants)
	review.POST("/update", handlers.ReviewUpdate)

	// start server with logger
	e.Logger.Fatal(e.Start(":42069"))
}
