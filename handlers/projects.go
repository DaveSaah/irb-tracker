package handlers

import (
	"net/http"

	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/boring-school-work/irb-tracker/types"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type projectsData struct {
	Session        *sessions.Session
	Error          string
	Projects       []model.Project
	UserInfo       model.User
	SubmittedCount int
	ApprovedCount  int
	RejectedCount  int
	PendingCount   int
}

// ProjectsView renders the projects page.
func ProjectsView(c echo.Context) error {
	sess, isLoggedIn := helpers.CheckSession(c)
	if !isLoggedIn {
		return c.Render(http.StatusBadRequest, "index", nil)
	}

	id := sess.Values["id"].(int)

	userinfo, err := functions.FetchUser(id)
	if err != nil {
		return err
	}

	submitted_count, err := functions.GetProjectCount(id, types.StatusSubmitted)
	if err != nil {
		return err
	}

	approved_count, err := functions.GetProjectCount(id, types.StatusApproved)
	if err != nil {
		return err
	}

	rejected_count, err := functions.GetProjectCount(id, types.StatusRejected)
	if err != nil {
		return err
	}

	pending_count, err := functions.GetProjectCount(id, types.StatusPending)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "projects",
		projectsData{
			Session:        sess,
			UserInfo:       userinfo,
			SubmittedCount: submitted_count,
			ApprovedCount:  approved_count,
			RejectedCount:  rejected_count,
			PendingCount:   pending_count,
		},
	)
}

// ProjectSearch will return a list of projects that match a
// specified pattern
func ProjectSearch(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)
	pattern := c.FormValue("search")

	id := sess.Values["id"].(int)

	if pattern == "" {
		return nil
	}

	data, err := functions.FindProjects(id, pattern)
	if err != nil {
		return c.Render(http.StatusBadRequest, "project_result",
			projectsData{Error: "Could not find any match"},
		)
	}

	return c.Render(http.StatusOK, "project_result", projectsData{Projects: data})
}
