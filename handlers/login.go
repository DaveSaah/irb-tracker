package handlers

import (
	"net/http"
	"strings"

	"github.com/boring-school-work/irb-tracker/actions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/labstack/echo/v4"
)

type loginData struct {
	Email  string
	Passwd string
	Error  string
}

// LoginView renders the login page
func LoginView(c echo.Context) error {
	sess, isLoggedIn := helpers.CheckSession(c)
	if isLoggedIn {
		return c.Render(http.StatusOK, "dashboard", sess)
	}

	return c.Render(http.StatusOK, "login", nil)
}

// LoginUser logs in a user
func LoginUser(c echo.Context) error {
	email := c.FormValue("email")
	passwd := c.FormValue("passwd")

	user, err := actions.LoginUser(email, passwd)
	if err != nil {
		if strings.Contains(err.Error(), "email") || strings.Contains(err.Error(), "password") {
			data := loginData{
				Email:  email,
				Passwd: passwd,
				Error:  "Incorrect email or password",
			}

			return c.Render(http.StatusBadRequest, "login", data)
		} else {
			return err
		}
	}

	sess := helpers.CreateSession(&user, c)
	return c.Render(http.StatusOK, "dashboard", sess)
}
