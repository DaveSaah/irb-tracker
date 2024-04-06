package helpers

import (
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// CreateSession creates a new session for the user
func CreateSession(u *model.User, c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		MaxAge:   86400, // 24 hours
		HttpOnly: true,
	}

	sess.Values["id"] = u.ID
	sess.Values["username"] = u.FName + " " + u.LName
	sess.Values["type"] = u.Type

	sess.Save(c.Request(), c.Response())
	return sess
}

// DestroySession destroys the session for the user
func DestroySession(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		MaxAge: -1,
	}

	return sess.Save(c.Request(), c.Response())
}

// CheckSession checks if the user is logged in or not
// returns the session and a boolean value
func CheckSession(c echo.Context) (*sessions.Session, bool) {
	sess, _ := session.Get("session", c)
	if sess.Values["id"] == nil {
		return nil, false
	}

	return sess, true
}
