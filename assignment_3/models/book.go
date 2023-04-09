package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint    `gorm:"primary_key" json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

type CreateBookInput struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

type UpdateBookInput struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}
