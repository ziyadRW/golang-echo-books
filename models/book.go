package models

import "gorm.io/gorm"

type Book struct {
	gorm.model

	Title string `json: "title"`
	Authot string `json: "author"`
	Published int `json: "publisehd"`
}