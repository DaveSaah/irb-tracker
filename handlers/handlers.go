package handlers

import (
	"net/http"

	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/labstack/echo/v4"
)

// IndexView renders the index page.
func IndexView(c echo.Context) error {
	_, isLoggedIn := helpers.CheckSession(c)
	if isLoggedIn {
		return DashboardView(c)
	}

	return c.Render(http.StatusOK, "index", nil)
}

// HomeView renders the home page.
func HomeView(c echo.Context) error {
	_, isLoggedIn := helpers.CheckSession(c)
	if isLoggedIn {
		return DashboardView(c)
	}

	return c.Render(http.StatusOK, "home", nil)
}

// ReviewView renders the review page.
func ReviewView(c echo.Context) error {
	sess, isLoggedIn := helpers.CheckSession(c)
	if !isLoggedIn {
		return c.Render(http.StatusBadRequest, "index", nil)
	}

	if sess.Values["type"] != "faculty" {
		return DashboardView(c)
	}

	return c.Render(http.StatusOK, "review", nil)
}

// LogoutUser logouts out a user
func LogoutUser(c echo.Context) error {
	helpers.DestroySession(c)
	return c.Render(http.StatusOK, "home", nil)
}
