package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DB_CONNECT")
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		DB = db
	}
}
