package utils

import (
	"errors"

	"github.com/ziyadRW/golang-echo-books/DTOs"
)

func ValidateBook(book *DTOs.Book) error {
	if book.Title == "" {
		return errors.New("title is required")
	}
	if book.Authotr == "" {
		return errors.New("author is required")
	}
	if book.Published <= 0 {
		return errors.New("published year must be greater than 0")
	}

	return nil
}