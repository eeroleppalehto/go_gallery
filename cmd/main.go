package main

import (
	"github.com/eeroleppalehto/go_gallery/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	app := echo.New()

	app.Static("/static", "static")

	app.Use(middleware.Logger())

	homeHandler := handler.HomeHandler{}
	galleryHandler := handler.GalleryHandler{}
	creatorshandler := handler.CreatorsHandler{}

	app.GET("/", homeHandler.HandleHomeShow)

	app.GET("/gallery", galleryHandler.HandleGalleryShow)

	app.GET("/photos/:imageID", galleryHandler.HandlePhotoShow)

	app.GET("/creators", creatorshandler.HandlePhotographerShow)

	app.Logger.Fatal(app.Start(":8081"))
}
