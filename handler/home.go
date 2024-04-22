package handler

import (
	"github.com/eeroleppalehto/go_gallery/views/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (r *RouteHandler) HomeShow(c echo.Context) error {

	return render(c, home.Show())
}
