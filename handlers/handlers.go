package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// IndexView renders the index page.
func IndexView(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// HomeView renders the home page.
func HomeView(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}

// DashboardView renders the dashboard page.
func DashboardView(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard", nil)
}

// ProjectsView renders the projects page.
func ProjectsView(c echo.Context) error {
	return c.Render(http.StatusOK, "projects", nil)
}

// ActivityView renders the activity page.
func ActivityView(c echo.Context) error {
	return c.Render(http.StatusOK, "activity", nil)
}

// ReviewView renders the review page.
func ReviewView(c echo.Context) error {
	return c.Render(http.StatusOK, "review", nil)
}
