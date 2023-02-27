package db

import (
	"api/docs/structs"
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

	err1 := db.AutoMigrate(&structs.Comment{}, &structs.Book{}, &structs.User{})
	if err1 != nil {
		return nil
	}

	return db
}
