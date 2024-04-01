package main

import (
	"time"

	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/boring-school-work/irb-tracker/tmpl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Department struct {
	Depts []model.Department
}

type StudentData struct {
	Majors       []model.Major
	YearGroupMax int
	YearGroupMin int
}

type Error struct {
	MailError   string
	PasswdError string
}

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

	e.GET("/register", func(c echo.Context) error {
		depts, err := functions.FetchDepts()
		if err != nil {
			return err
		}

		return c.Render(200, "register", &Department{Depts: depts})
	})

	e.POST("/register/email", func(c echo.Context) error {
		email := c.FormValue("email")

		if ok := helpers.IsMailValid(email); !ok {
			return c.Render(200, "mail_error", &Error{MailError: "Invalid email"})
		}

		exists, err := functions.DoesMailExist(email)
		if err != nil {
			return err
		}

		if exists {
			return c.Render(200, "mail_error",
				&Error{MailError: "Email already exists"},
			)
		}

		return nil
	})

	e.POST("/register/student", func(c echo.Context) error {
		majors, err := functions.FetchMajors()
		if err != nil {
			return err
		}

		user_type := c.FormValue("type")

		if user_type == "student" {
			return c.Render(200, "student_info",
				&StudentData{
					Majors:       majors,
					YearGroupMin: time.Now().Year(),
					YearGroupMax: time.Now().Year() + 4,
				})
		}

		return nil
	})

	e.POST("/register/passwd", func(c echo.Context) error {
		passwd1 := c.FormValue("passwd1")
		passwd2 := c.FormValue("passwd2")

		if passwd2 != "" && passwd1 != passwd2 {
			return c.Render(200, "passwd_error",
				&Error{PasswdError: "Passwords do not match"},
			)
		}

		return nil
	})

	e.POST("/register", func(c echo.Context) error {
		return nil
	})

	// start server with logger
	e.Logger.Fatal(e.Start(":42069"))
}
