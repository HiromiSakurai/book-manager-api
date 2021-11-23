package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `gorm:"size:255"`
}
