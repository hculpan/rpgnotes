package models

import (
	"fmt"

	"gorm.io/gorm"
)

var Categories []Category

type Category struct {
	gorm.Model
	Name string
}

func FindCategory(name string) (*Category, error) {
	for _, cat := range Categories {
		if cat.Name == name {
			return &cat, nil
		}
	}

	return nil, fmt.Errorf("category '%s' not found", name)
}

func FindCategoryByID(id uint) (*Category, error) {
	for _, cat := range Categories {
		if cat.ID == id {
			return &cat, nil
		}
	}

	return nil, fmt.Errorf("category with id '%d' not found", id)
}
