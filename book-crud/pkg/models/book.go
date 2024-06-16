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
