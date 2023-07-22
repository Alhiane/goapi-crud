package models

import (
	"github.com/Alhiane/goapi-crud/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

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

// InitDB initializes the database
func init() {
	config.Connect()
	db = config.DB

	// run the migrations
	config.DB.AutoMigrate(&Book{})
}

// TableName sets the table name to plural
func (b *Book) TableName() string {
	return "users"
}

// create a new book
func (b *Book) Create() (*Book, error) {
	err := db.Create(&b).Error
	if err != nil {
		return nil, err
	}

	return b, nil
}

// return all books
func GetBooks() ([]*Book, error) {
	var books []*Book
	err := db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

// return a single book based on the id
func GetBook(id int) (*Book, error) {
	var book Book
	err := db.First(&book, id).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// update a book
func (b *Book) Update(id int) (*Book, error) {
	err := db.Save(&b).Error
	if err != nil {
		return nil, err
	}

	return b, nil
}

// delete a book
func DeleteBook(id int) error {
	var book Book
	err := db.Delete(&book, id).Error
	if err != nil {
		return err
	}

	return nil
}
