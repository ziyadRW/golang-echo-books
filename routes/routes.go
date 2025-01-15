package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/ziyadRW/golang-echo-books/handlers"
)

func SetUpRoutes(g *echo.Group) {
	
	bookHandler := handlers.NewBookHandler()
	
	g.GET("/books", bookHandler.GetAllBooks)
	g.GET("/book/ :id", bookHandler.GetBook)
	g.POST("/book", bookHandler.CreateBook)
	g.DELETE("/book/:id", bookHandler.DeleteBook)
	g.PUT("/book/ :id", bookHandler.UpdateBook)

}