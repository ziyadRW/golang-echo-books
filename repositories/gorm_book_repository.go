package repositories

import (
	"github.com/ziyadRW/golang-echo-books/DTOs"
	"gorm.io/gorm"
)

type GormBookRepository struct {
	DB *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository{
	return &GormBookRepository{DB: db}
}

func (r *GormBookRepository) CreateBook(book *DTOs.Book) error {
	return r.DB.Create(book).Error 
} 

// todo