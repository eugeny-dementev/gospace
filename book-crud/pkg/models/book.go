package models

import (
	"fmt"

	"example.com/book-crud/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Book) String() string {
	return fmt.Sprintf("(%v): %v - %v", b.ID, b.Author, b.Name)
}

func init() {
  config.Connect()
  db = config.GetDB()
  db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
  db.NewRecord(b)
  db.Create(&b)

  return b
}

func GetAllBooks() []Book {
  var Books []Book
  db.Find(&Books)
  return Books
}

func GetBookById(id uint) (*Book, *gorm.DB) {
  var book Book
  db := db.Where("ID = ?", id).Find(&book)
  return &book, db
}

func DeleteBook(id uint) Book {
  var book Book
  db.Where("ID = ?", id).Delete(&book)
  return book
}
