package main

import (
	"net/http"
	"https://github.com/ziyadRW/golang-echo-books/models"
	"https://github.com/ziyadRW/golang-echo-books/routes"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"

)

func main() {
	// start echo
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.use(middleware.Recover())


	//database 
	db := models.InitDB("books.db")
	models.Migrate(db)

	//routes
	api := e.Group("/api")
	routes.SetUpRoutes(api)

	// choose port
	e.Logger.Fatal(e.Start(":8080"))
}