package authservice

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CreateSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return err
}

func ReadSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	return err
}
