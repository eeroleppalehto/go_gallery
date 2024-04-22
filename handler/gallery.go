package handler

import (
	"strconv"

	"github.com/eeroleppalehto/go_gallery/views/gallery"
	"github.com/eeroleppalehto/go_gallery/views/photo"
	"github.com/labstack/echo/v4"
)

func (r *RouteHandler) GalleryShow(c echo.Context) error {
	images, err := r.Queries.GetPhotos(c.Request().Context())
	if err != nil {
		return c.String(404, "Failed to fetch photos")
	}

	return render(c, gallery.Show(images))
}

func (r *RouteHandler) PhotoShow(c echo.Context) error {
	imageIdStr := c.Param("imageID")

	imageIdInt, err := strconv.Atoi(imageIdStr)
	if err != nil {
		return c.String(400, "Bad request")
	}

	photograph, err := r.Queries.GetPhoto(c.Request().Context(), uint32(imageIdInt))
	if err != nil {
		return c.String(404, "Image not found")
	}

	return render(c, photo.Show(photograph))
}
