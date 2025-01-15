package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ziyadRW/golang-echo-books/models"
	"github.com/ziyadRW/golang-echo-books/repositories"
)

type Bookhandler struct {
	Repo repositories.BookReoisitory
}

func NewBookHandler(c echo.Context) *Bookhandler {
	return &Bookhandler{}
}

func (h *Bookhandler)GetAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "List of books will be returend here",
	})
}

func (h *Bookhandler)CreateBook(c echo.Context) error {
	book := new(models.Book)
	if err:= c.Bind(book); err != nil {
		return err
	}
	
	return nil
}

func (h *Bookhandler)GetBook(c echo.Context) error {
	return nil
}

func (h *Bookhandler)DeleteBook(c echo.Context) error {
	return nil
}

func (h *Bookhandler)UpdateBook(c echo.Context) error {
	return nil
}
