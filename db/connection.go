package db

import (
	"github.com/jinzhu/gorm"
	"pricecrawl/models"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=pricecrawl.tech sslmode=disable")

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