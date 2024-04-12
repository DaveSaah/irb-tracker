package handlers

import (
	"net/http"
	"strconv"

	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/labstack/echo/v4"
)

type projectData struct {
	Project model.Project
}

func ProjectView(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchProject(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "project_item",
		projectData{
			Project: project,
		},
	)
}

func ProjectInfo(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchProject(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "project_content",
		projectData{
			Project: project,
		},
	)
}

func ProjectMsg(c echo.Context) error {
	return c.Render(http.StatusOK, "messages", nil)
}

func ProjectBrief(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchProject(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "project_brief",
		projectData{
			Project: project,
		},
	)
}

func ProjectTimeline(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchProject(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "timeline",
		projectData{
			Project: project,
		},
	)
}

func ProjectParticipants(c echo.Context) error {
	sess, _ := helpers.CheckSession(c)

	_projID := c.QueryParam("id")
	projID, _ := strconv.Atoi(_projID)
	userID := sess.Values["id"].(int)

	project, err := functions.FetchProject(userID, projID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "participants_info",
		projectData{
			Project: project,
		},
	)
}
