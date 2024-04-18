package handler

import (
	"github.com/eeroleppalehto/go_gallery/models"
	"github.com/eeroleppalehto/go_gallery/views/creators"
	"github.com/labstack/echo/v4"
)

type CreatorsHandler struct{}

func (h *CreatorsHandler) HandlePhotographerShow(c echo.Context) error {

	user := models.User{
		ID:       1,
		Username: "Sabber",
		Email:    "",
		Password: "",
	}

	return render(c, creators.Show(user))
}
