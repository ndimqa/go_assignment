package db

import (
	"api/docs/books"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/api"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err2 := db.AutoMigrate(&books.Book{})
	if err2 != nil {
		log.Fatal(err2)
		return nil
	}

	return db
}
