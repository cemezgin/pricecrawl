package db

import (
	"crawltest/models"
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=crawltest sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	AutoMigrate(db)

	//defer db.Close()

	return db
}

func AutoMigrate(db *gorm.DB) {
	objects := []interface{}{
		&models.Site{}, &models.Product{},
	}

	for _, item := range objects {
		if item == nil {
			continue
		}
		db.AutoMigrate(item)
	}
}