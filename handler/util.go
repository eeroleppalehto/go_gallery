package handler

import (
	"github.com/a-h/templ"
	"github.com/eeroleppalehto/go_gallery/models"
	authservice "github.com/eeroleppalehto/go_gallery/service/authService"
	"github.com/eeroleppalehto/go_gallery/views/layout"
	"github.com/labstack/echo/v4"
)

type RouteHandler struct {
	Queries  *models.Queries
	Sessions *authservice.SessionService
}

func render(c echo.Context, component templ.Component) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return component.Render(c.Request().Context(), c.Response())
	}

	auth := r.Sessions.IsAuthenticated(c)

	return layout.Base(component, auth.IsAuthenticated, auth.Username).Render(c.Request().Context(), c.Response())
}

// func NotFound(c echo.Context) error {
// 	return r.render(c, notFound.Show())
// }
