package handler

import (
	"github.com/eeroleppalehto/go_gallery/models"
	"github.com/eeroleppalehto/go_gallery/views/gallery"
	"github.com/labstack/echo/v4"
)

type GalleryHandler struct {
}

func (h *GalleryHandler) HandleGalleryShow(c echo.Context) error {
	images := models.GetImages()

	return render(c, gallery.Show(images))
}
