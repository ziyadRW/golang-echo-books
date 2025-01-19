package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/ziyadRW/golang-echo-books/DTOs"
	"github.com/ziyadRW/golang-echo-books/repositories"
	"github.com/ziyadRW/golang-echo-books/utils"
)

type Bookhandler struct {
	Repo repositories.BookReoisitory
}

func NewBookHandler(repo repositories.BookReoisitory) *Bookhandler {
	return &Bookhandler{Repo: repo}
}

func (h *Bookhandler)GetAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "List of books will be returend here",
	})
}

func (h *Bookhandler)CreateBook(c echo.Context) error {
	book := new(DTOs.Book)
	if err:= c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//validate rquest
	validated := utils.ValidateBook(book)
	if validated != nil {
		log.Error(validated.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create book")
	}
	//create book
	err := h.Repo.CreateBook(book)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create book")
	}
	
	return c.JSON(http.StatusCreated, book)
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
