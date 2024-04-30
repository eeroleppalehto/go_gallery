package main

import (
	"database/sql"
	"log"

	"github.com/eeroleppalehto/go_gallery/handler"
	"github.com/eeroleppalehto/go_gallery/models"
	authservice "github.com/eeroleppalehto/go_gallery/service/authService"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading variables from .env: ", err)
	}

	app := echo.New()

	app.Static("/static", "static")

	app.Use(middleware.Logger())

	queries, err := getQueryEngine()
	if err != nil {
		log.Fatal(err)
	}

	routeHandler := handler.RouteHandler{
		Queries:  queries,
		Sessions: &authservice.SessionService{},
	}

	routeHandler.Sessions.Init()

	app.GET("/", routeHandler.HomeShow)

	app.GET("/gallery", routeHandler.GalleryShow)

	app.GET("/photos/:imageID", routeHandler.PhotoShow)

	app.GET("/creators", routeHandler.PhotographerShow)

	app.GET("/login", routeHandler.LoginForm)
	app.POST("login", routeHandler.Login)

	app.POST("/logout", routeHandler.Logout)

	app.GET("/sign-up", routeHandler.SignupShow)
	app.POST("/sign-up", routeHandler.Signup)

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
