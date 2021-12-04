package model

import (
	"time"
)

type Book struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"size:255" gorm:"not null" json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
