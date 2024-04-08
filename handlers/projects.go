package handlers

import (
	"net/http"

	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/labstack/echo/v4"
)

type projectData struct {
	Error    string
	Projects []model.Project
}

// ProjectsView renders the projects page.
func ProjectsView(c echo.Context) error {
	_, isLoggedIn := helpers.CheckSession(c)
	if !isLoggedIn {
		return c.Render(http.StatusBadRequest, "index", nil)
	}

	return c.Render(http.StatusOK, "projects", nil)
}

// ProjectSearch will return a list of projects that match a
// specified pattern
func ProjectSearch(c echo.Context) error {
	pattern := c.FormValue("search")

	if pattern == "" {
		return nil
	}

	data, err := functions.FindProjects(pattern)
	if err != nil {
		return c.Render(http.StatusBadRequest, "project_result",
			projectData{Error: "Could not find any match"},
		)
	}

	return c.Render(http.StatusOK, "project_result", projectData{Projects: data})
}
