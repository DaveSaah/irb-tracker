package handlers

import (
	"net/http"

	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type msgData struct {
	Session *sessions.Session
	Status  []model.Status
}

func ProjectMsg(c echo.Context) error {
	return c.Render(http.StatusOK, "messages", nil)
}

func ReviewMsg(c echo.Context) error {
	status, err := functions.FetchStatus()
	if err != nil {
		return err
	}

	sess, _ := helpers.CheckSession(c)

	return c.Render(http.StatusOK, "messages",
		msgData{
			Status:  status,
			Session: sess,
		},
	)
}
