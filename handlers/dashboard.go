package handlers

import (
	"net/http"

	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type dashboardData struct {
	Session         *sessions.Session
	Proposals       []model.Project
	PendingPercent  int
	ApprovedPercent int
	RejectedPercent int
	SubmittedCount  int
	ApprovedCount   int
	RejectedCount   int
}

// DashboardView renders the dashboard page.
func DashboardView(c echo.Context) error {
	sess, isLoggedIn := helpers.CheckSession(c)
	if !isLoggedIn {
		return c.Render(http.StatusBadRequest, "index", nil)
	}

	id := sess.Values["id"].(int)

	submitted_count, err := functions.GetProjectCount(id, functions.StatusSubmitted)
	if err != nil {
		return err
	}

	approved_count, err := functions.GetProjectCount(id, functions.StatusApproved)
	if err != nil {
		return err
	}

	rejected_count, err := functions.GetProjectCount(id, functions.StatusRejected)
	if err != nil {
		return err
	}

	approved_percent := (approved_count / submitted_count) * 100
	rejected_percent := (rejected_count / submitted_count) * 100
	pending_percent := 100 - approved_percent - rejected_percent

	proposals, err := functions.FetchProjects(id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "dashboard",
		dashboardData{
			PendingPercent:  pending_percent,
			ApprovedPercent: approved_percent,
			RejectedPercent: rejected_percent,
			Session:         sess,
			SubmittedCount:  submitted_count,
			ApprovedCount:   approved_percent,
			RejectedCount:   rejected_count,
			Proposals:       proposals,
		},
	)
}
