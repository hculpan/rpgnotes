package config

import "github.com/hculpan/rpgnotes/models"

func LoadCateogires() {
	result := DB.Find(&models.Categories)
	if result.Error != nil {
		panic(result.Error)
	}
}
