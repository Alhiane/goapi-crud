package models

import (
	"gorm.io/gorm"
)

// Book is a representation of a book
type Book struct {
	gorm.Model
	Title       string  `gorm:"type:varchar(100)" json:"title"`
	Author      string  `gorm:"type:varchar(100)" json:"author"`
	Publisher   string  `gorm:"type:varchar(100)" json:"publisher"`
	PublishDate string  `gorm:"type:date" json:"publish_date"`
	ISBN        string  `gorm:"type:varchar(13);unique" json:"isbn"`
	Genre       string  `gorm:"type:varchar(50)" json:"genre"`
	Pages       int     `gorm:"type:int" json:"pages"`
	Price       float64 `gorm:"type:decimal(5,2)" json:"price"`
	Language    string  `gorm:"type:varchar(20)" json:"language"`
	Description string  `gorm:"type:text" json:"description"`
}
