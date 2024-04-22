package handler

import (
	"github.com/a-h/templ"
	"github.com/eeroleppalehto/go_gallery/models"
	"github.com/eeroleppalehto/go_gallery/views/layout"
	"github.com/eeroleppalehto/go_gallery/views/notFound"
	"github.com/labstack/echo/v4"
)

type RouteHandler struct {
	Queries *models.Queries
}

func render(c echo.Context, component templ.Component) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return component.Render(c.Request().Context(), c.Response())
	}
	return layout.Base(component).Render(c.Request().Context(), c.Response())
}

func NotFound(c echo.Context) error {
	return render(c, notFound.Show())
}
