package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title		string `json:"title"`
	Authotr 		string `json:"author"`
	Published 	int `json:"publisehd"`
}