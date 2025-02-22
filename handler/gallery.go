package handler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	s "strings"

	"github.com/eeroleppalehto/go_gallery/views/gallery"
	"github.com/eeroleppalehto/go_gallery/views/photo"
	"github.com/labstack/echo/v4"
)

func (r *RouteHandler) GalleryShow(c echo.Context) error {
	queries := getQueryEngine(r.DB)
	images, err := queries.GetPhotos(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return c.String(404, "Failed to fetch photos")
	}

	return r.render(c, gallery.Show(images))
}

func (r *RouteHandler) PhotoShow(c echo.Context) error {
	imageIdStr := c.Param("imageID")

	imageIdInt, err := strconv.Atoi(imageIdStr)
	if err != nil {
		return c.String(400, "Bad request")
	}

	queries := getQueryEngine(r.DB)

	photograph, err := queries.GetPhoto(c.Request().Context(), uint32(imageIdInt))
	if err != nil {
		return c.String(404, "Image not found")
	}

	return r.render(c, photo.Show(photograph, false))
}

func (r *RouteHandler) PhotoDelete(c echo.Context) error {
	imageIdStr := c.Param("imageID")

	imageIdInt, err := strconv.Atoi(imageIdStr)
	if err != nil {
		return c.String(400, "Bad request")
	}

	queries := getQueryEngine(r.DB)

	photograph, err := queries.GetPhoto(c.Request().Context(), uint32(imageIdInt))
	if err != nil {
		return c.String(404, "Image not found")
	}

	_, err = queries.DeletePhoto(c.Request().Context(), uint32(imageIdInt))
	if err != nil {
		return c.String(404, "Image not found")
	}

	err = os.Remove(fmt.Sprintf("static/images/%s", photograph.Filename))
	if err != nil {
		log.Fatal(err)
		fmt.Println()
	}

	err = os.Remove(fmt.Sprintf("static/images-lq/%s", s.Replace(photograph.Filename, ".jpg", "-lq.jpg", -1)))
	if err != nil {
		log.Fatal(err)
	}

	return r.render(c, photo.Show(photograph, true))
}
