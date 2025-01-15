package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ziyadRW/golang-echo-books/models"
	"github.com/ziyadRW/golang-echo-books/routes"
)


func main() {
	// start echo
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	//database 
	db := models.InitDB("books.db")
	models.Migrate(db)

	//routes
	api := e.Group("/api")
	routes.SetUpRoutes(api)

	// choose port
	e.Logger.Fatal(e.Start(":8080"))
}