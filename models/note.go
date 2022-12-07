package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title      string
	Note       string
	CategoryID uint
	Keywords   string
	Category   Category
}
