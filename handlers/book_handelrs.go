package bookHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	return nil
}

func Book(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "List of books will be returend here",
	})
}