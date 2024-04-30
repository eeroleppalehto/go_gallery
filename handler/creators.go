package handler

import (
	"github.com/eeroleppalehto/go_gallery/views/creators"
	"github.com/labstack/echo/v4"
)

func (r *RouteHandler) PhotographerShow(c echo.Context) error {
	queries := getQueryEngine(r.DB)
	users, err := queries.GetUsers(c.Request().Context())
	if err != nil {
		return c.String(404, "Failed to fetch users")
	}

	return r.render(c, creators.Show(users[0]))
}
