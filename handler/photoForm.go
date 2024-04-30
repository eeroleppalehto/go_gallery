package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	imageservice "github.com/eeroleppalehto/go_gallery/service/imageService"
	"github.com/eeroleppalehto/go_gallery/views/photo"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (r *RouteHandler) PhotoForm(c echo.Context) error {
	return r.render(c, photo.Form())
}

func (r *RouteHandler) PostPhoto(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")

	// Upload file
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to upluad file")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	uuid := uuid.New()
	// Destination
	pathNorm := "static/images/"
	pathLQ := "static/images-lq/"
	fileName := "{user}-{uuid}"
	fileExtension := "jpg"

	imgServ := imageservice.ImageService{}
	imgServ.JPG.LoadImageM(src)

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// TODO: Save photo info to db (make it transactional and also generate the needed fields), generate a LQ version of the image

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields title=%s and desc=%s.</p>", file.Filename, title, description))
}
