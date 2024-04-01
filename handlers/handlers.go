package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index renders the index page.
func IndexView(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// Home renders the home page.
func HomeView(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
