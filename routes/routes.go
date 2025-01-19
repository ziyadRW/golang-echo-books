package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/ziyadRW/golang-echo-books/handlers"
	"github.com/ziyadRW/golang-echo-books/models"
	"github.com/ziyadRW/golang-echo-books/repositories"
)

func SetUpRoutes(g *echo.Group) {
	
	repo := repositories.NewGormBookRepository(models.DB)
	bookHandler := handlers.NewBookHandler(repo)
	
	g.GET("/books", bookHandler.GetAllBooks)
	g.GET("/book/ :id", bookHandler.GetBook)
	g.POST("/book", bookHandler.CreateBook)
	g.DELETE("/book/:id", bookHandler.DeleteBook)
	g.PUT("/book/ :id", bookHandler.UpdateBook)

}