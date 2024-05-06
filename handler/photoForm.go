package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/eeroleppalehto/go_gallery/models"
	imageservice "github.com/eeroleppalehto/go_gallery/service/imageService"
	"github.com/eeroleppalehto/go_gallery/views/photo"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (r *RouteHandler) PhotoForm(c echo.Context) error {
	return r.render(c, photo.Form())
}

func (r *RouteHandler) PostPhoto(c echo.Context) error {
	// Begin database transaction
	tx, err := r.DB.Begin()
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to start db tx")
	}
	defer tx.Rollback()

	queries := getQueryEngine(r.DB)

	qtx := queries.WithTx(tx)

	// Read fields from the request and and upload the jpg file
	title := c.FormValue("title")
	description := c.FormValue("description")

	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to upload the image")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to upload the image")
	}
	defer src.Close()

	// Add a row to the photo table
	username := r.Sessions.IsAuthenticated(c).Username
	user, err := queries.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		return err
	}
	filename, filenameLowQuality := generateFileNames(username)

	_, err = qtx.CreatePhoto(c.Request().Context(), models.CreatePhotoParams{
		UserID:      user.UserID,
		Title:       title,
		Description: sql.NullString{String: description, Valid: true},
		Filename:    filename,
		Date:        time.Now(),
	})
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "Error while storing to DB")
	}

	path := "static/images/"
	pathLQ := "static/images-lq/"

	// Decode the jpg-file, generate a low quality version of the image and store them
	imgServ := imageservice.ImageService{}
	image, err := imgServ.JPG.LoadImageFromReader(src)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "Error while processing image 1")
	}

	point, err := imgServ.GetNewBounds(image, 500_000)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "Error while processing image 2")
	}

	newImg, err := imgServ.Resize(image, point)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "Error while processing image 3")
	}

	err = imgServ.JPG.SaveImage(newImg, pathLQ+filenameLowQuality)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "Error while processing image 4"+err.Error())
	}

	err = imgServ.JPG.SaveImage(image, path+filename)
	if err != nil {
		os.Remove(pathLQ + filenameLowQuality)
		return c.HTML(http.StatusInternalServerError, "Error while processing image 5"+err.Error())
	}

	// TODO: If error occurs, destroy any images that has been stored

	err = tx.Commit()
	if err != nil {
		os.Remove(pathLQ + filenameLowQuality)
		os.Remove(pathLQ + filename)
		return c.HTML(http.StatusInternalServerError, "Error while commiting to DB")
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields title=%s and desc=%s.</p>", file.Filename, title, description))
}

func generateFileNames(username string) (filename string, filenameLowQuality string) {
	uuid := uuid.New()
	filename = "gollery" + username + uuid.String() + ".jpg"
	filenameLowQuality = "gollery" + username + uuid.String() + "-lq.jpg"

	return filename, filenameLowQuality
}
