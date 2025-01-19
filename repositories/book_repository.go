package repositories

import (
	"github.com/ziyadRW/golang-echo-books/DTOs"
	"github.com/ziyadRW/golang-echo-books/models"
)

type BookReoisitory interface {
	CreateBook( book *DTOs.Book) error
	GetAllBooks() ( []models.Book, error)
	GetBookById(id uint) ( *models.Book, error)
	UpdateBook(id uint, book *DTOs.Book) ( *models.Book, error)
	DeleteBook( id uint) error

}