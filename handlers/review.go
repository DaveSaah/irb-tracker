package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/boring-school-work/irb-tracker/types"
	"github.com/labstack/echo/v4"
)

type reviewData struct {
	Error         string
	Projects      []model.Project
	ApprovedCount int
	RejectedCount int
	PendingCount  int
}

// ReviewView renders the review page.
func ReviewView(c echo.Context) error {
	sess, isLoggedIn := helpers.CheckSession(c)
	if !isLoggedIn {
		return c.Render(http.StatusBadRequest, "index", nil)
	}

	if sess.Values["type"] != "faculty" {
		return c.Redirect(http.StatusSeeOther, "/dashboard")
	}

	id := sess.Values["id"].(int)
	data, err := functions.FetchAllProjectsForReview(id)
	if err != nil {
		return err
	}

	approved_count, err := functions.GetProjectReviewCount(id, types.StatusApproved)
	if err != nil {
		return err
	}

	rejected_count, err := functions.GetProjectReviewCount(id, types.StatusRejected)
	if err != nil {
		return err
	}

	pending_count, err := functions.GetProjectReviewCount(id, types.StatusPending)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "review",
		reviewData{
			Projects:      data,
			ApprovedCount: approved_count,
			RejectedCount: rejected_count,
			PendingCount:  pending_count,
		},
	)
}

func ReviewSearch(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)
	pattern := c.FormValue("search")

	id := sess.Values["id"].(int)

	if pattern == "" {
		return nil
	}

	data, err := functions.FindProjectForReview(id, pattern)
	if err != nil {
		return c.Render(http.StatusBadRequest, "project_result",
			reviewData{Error: "Could not find any match"},
		)
	}

	return c.Render(http.StatusOK, "review_result", reviewData{Projects: data})
}

func ProjectReview(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchReviewItem(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "review_item",
		projectData{
			Project: project,
		},
	)
}

func ReviewInfo(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchReviewItem(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "project_review_content",
		projectData{
			Project: project,
		},
	)
}

func ReviewBrief(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchReviewItem(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "project_review_brief",
		projectData{
			Project: project,
		},
	)
}

func ReviewTimeline(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchReviewItem(userID, projID)
	if err != nil {
		return err
	}

	log.Printf("Date submitted: %s", project.DateSubmitted)

	return c.Render(http.StatusOK, "review_timeline",
		projectData{
			Project: project,
		},
	)
}

func ReviewParticipants(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchReviewItem(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "review_participants_info",
		projectData{
			Project: project,
		},
	)
}
