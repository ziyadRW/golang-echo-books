package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/ziyadRW/golang-echo-books/bookHandler"
)

func SetUpRoutes(g *echo.Group) {
	
	g.GET("/books", bookHandler.Book)

}