package model

import (
	"time"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:255" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
