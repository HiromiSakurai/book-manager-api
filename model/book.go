package model

import (
	"time"

	"github.com/HiromiSakurai/book-manager-api/database"
)

type Book struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"size:255" gorm:"not null" json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (book *Book) Get() error {
	result := database.DB.First(book)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (book *Book) Save() error {
	result := database.DB.Create(book)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
