package handler

import (
	"net/http"

	"github.com/eeroleppalehto/go_gallery/views/login"

	"github.com/labstack/echo/v4"
)

func (r *RouteHandler) LoginForm(c echo.Context) error {
	form := login.NewLoginForm()

	authState := r.Sessions.IsAuthenticated(c)

	if authState.IsAuthenticated {
		form.IsSuccess = true
	}

	return r.render(c, login.Form(form))
}

func (r *RouteHandler) Login(c echo.Context) error {
	form := login.LoginForm{
		Username:   c.FormValue("username"),
		Password:   c.FormValue("password"),
		IsSuccess:  false,
		LoginError: false,
	}

	status := r.Sessions.Login(c.Request().Context(), c.Request(), c.Response(), r.DB)
	c.Response().Status = status
	if status != http.StatusOK {
		form.LoginError = true
		return r.render(c, login.Form(form))
	}

	form.IsSuccess = true

	return r.render(c, login.Form(form))
}

func (r *RouteHandler) Logout(c echo.Context) error {
	err := r.Sessions.Logout(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return r.render(c, login.Logout())
}
