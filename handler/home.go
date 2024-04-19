package handler

import (
	"fmt"

	"github.com/eeroleppalehto/go_gallery/views/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (h *HomeHandler) HandleHomeShow(c echo.Context) error {
	cc := c.(*CustomContext)
	users, err := cc.Queries.GetUsers(c.Request().Context())
	if err != nil {
		return err
	}

	fmt.Println(users[0].Username)

	return render(c, home.Show())
}
