package main

import (
	"github.com/hculpan/rpgnotes/config"
	"github.com/hculpan/rpgnotes/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func addCategories() error {
	if err := addCategory("NPC"); err != nil {
		return err
	}
	if err := addCategory("Location"); err != nil {
		return err
	}
	if err := addCategory("Gear"); err != nil {
		return err
	}
	if err := addCategory("Plot"); err != nil {
		return err
	}

	return nil
}

func addCategory(name string) error {
	cat := models.Category{Name: name}
	result := config.DB.Create(&cat)
	return result.Error
}

func main() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Category{})
	config.DB.AutoMigrate(&models.Note{})

	result := config.DB.Exec("DELETE FROM notes")
	if result.Error != nil {
		panic(result.Error)
	}

	result = config.DB.Exec("DELETE FROM categories")
	if result.Error != nil {
		panic(result.Error)
	}

	if err := addCategories(); err != nil {
		panic(err)
	}
}
