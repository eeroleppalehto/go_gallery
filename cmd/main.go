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

func main() {

	app := echo.New()

	app.Static("/static", "static")

	app.Use(middleware.Logger())

	queries, err := getQueryEngine()
	if err != nil {
		log.Fatal(err)
	}

	routeHandler := handler.RouteHandler{
		Queries: queries,
	}

	app.GET("/", routeHandler.HomeShow)

	app.GET("/gallery", routeHandler.GalleryShow)

	app.GET("/photos/:imageID", routeHandler.PhotoShow)

	app.GET("/creators", routeHandler.PhotographerShow)

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
