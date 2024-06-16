package models

import "fmt"

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	ID     uint32 `json:"id"`
}

func (b *Book) String() string {
	return fmt.Sprintf("(%v): %v - %v", b.ID, b.Author, b.Name)
}
