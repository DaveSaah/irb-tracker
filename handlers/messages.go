package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/boring-school-work/irb-tracker/actions"
	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type msgData struct {
	Session *sessions.Session
	Status  []model.Status
	ID      int
}

func ProjectMsg(c echo.Context) error {
	return c.Render(http.StatusOK, "messages", nil)
}

func ReviewMsg(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	status, err := functions.FetchStatus()
	if err != nil {
		return err
	}

	sess, _ := helpers.CheckSession(c)

	return c.Render(http.StatusOK, "messages",
		msgData{
			Status:  status,
			Session: sess,
			ID:      id,
		},
	)
}

func ReviewUpdate(c echo.Context) error {
	projID, _ := strconv.Atoi(c.QueryParam("id"))
	status_id, _ := strconv.Atoi(c.FormValue("status"))

	log.Println("Updating status for project:", c.QueryParams(), "to:", status_id)

	err := actions.UpdateStatus(projID, status_id)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/review")
}
