package main

import (
	"database/sql"
	"log"

	"github.com/eeroleppalehto/go_gallery/handler"
	"github.com/eeroleppalehto/go_gallery/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type CustomContext struct {
// 	echo.Context
// 	*models.Queries
// }

func main() {

	app := echo.New()

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		queries, err := getQueryEngine()
		if err != nil {
			log.Fatal(err)
		}
		return func(c echo.Context) error {
			cc := handler.CustomContext{
				Context: c,
				Queries: queries,
			}
			return next(cc)
		}
	})

	app.Static("/static", "static")

	app.Use(middleware.Logger())

	_, err := getQueryEngine()
	if err != nil {
		log.Fatal(err)
	}

	homeHandler := handler.HomeHandler{}
	galleryHandler := handler.GalleryHandler{}
	creatorshandler := handler.CreatorsHandler{}

	app.GET("/", homeHandler.HandleHomeShow)

	app.GET("/gallery", galleryHandler.HandleGalleryShow)

	app.GET("/photos/:imageID", galleryHandler.HandlePhotoShow)

	app.GET("/creators", creatorshandler.HandlePhotographerShow)

	app.Logger.Fatal(app.Start(":8081"))
}

func getQueryEngine() (*models.Queries, error) {
	db, err := sql.Open("mysql", "root:Q2werty@/gollery?parseTime=true")
	if err != nil {
		return nil, err
	}

	queries := models.New(db)

	return queries, nil
}
