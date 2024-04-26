package handler

import (
	"github.com/eeroleppalehto/go_gallery/views/creators"
	"github.com/labstack/echo/v4"
)

func (r *RouteHandler) PhotographerShow(c echo.Context) error {
	users, err := r.Queries.GetUsers(c.Request().Context())
	if err != nil {
		return c.String(404, "Failed to fetch users")
	}

	return render(c, creators.Show(users[0]))
}
