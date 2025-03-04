package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/eeroleppalehto/go_gallery/handler"
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

	connectionString := os.Getenv("DATABASE_URL")

	if os.Getenv("ENV") == "DEV" {
		connectionString = os.Getenv("DATABASE_URL_DEV")
	}

	app := echo.New()

	app.Static("/static", "static")

	app.Use(middleware.Logger())

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	routeHandler := handler.RouteHandler{
		DB:       db,
		Sessions: &authservice.SessionService{},
	}

	routeHandler.Sessions.Init()

	app.GET("/", routeHandler.HomeShow)

	app.GET("/gallery", routeHandler.GalleryShow)

	app.GET("/photos/:imageID", routeHandler.PhotoShow)
	app.GET("/photos/add-new", routeHandler.PhotoForm)
	app.POST("/photos/add-new", routeHandler.PostPhoto)
	app.DELETE("/photos/:imageID", routeHandler.PhotoDelete)

	app.GET("/creators", routeHandler.PhotographerShow)

	app.GET("/login", routeHandler.LoginForm)
	app.POST("login", routeHandler.Login)

	app.POST("/logout", routeHandler.Logout)

	app.GET("/sign-up", routeHandler.SignupShow)
	app.POST("/sign-up", routeHandler.Signup)

	app.Logger.Fatal(app.Start(":8081"))
}
