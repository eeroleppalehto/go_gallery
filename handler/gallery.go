package handler

import (
	"github.com/eeroleppalehto/go_gallery/views/gallery"
	"github.com/eeroleppalehto/go_gallery/views/photo"
	"github.com/labstack/echo/v4"
)

type GalleryHandler struct {
}

func (h *GalleryHandler) HandleGalleryShow(c echo.Context) error {
	// images := models.GetImages()

	return render(c, gallery.Show(images))
}

func (h *GalleryHandler) HandlePhotoShow(c echo.Context) error {
	imageID := c.Param("imageID")

	// image, err := models.GetImage(imageID)

	if err != nil {
		return c.String(404, "Image not found")
	}

	return render(c, photo.Show(image))
}
