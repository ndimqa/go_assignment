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

	//err1 := db.AutoMigrate(&structs.Comment{})
	//if err1 != nil {
	//	return nil
	//}

	err2 := db.AutoMigrate(&structs.Book{})
	if err2 != nil {
		log.Fatal(err2)
		return nil
	}

	err3 := db.AutoMigrate(&structs.User{})
	if err3 != nil {
		return nil
	}

	return db
}
