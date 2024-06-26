package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type URls struct {
	gorm.Model
	URL string
	Key string
}

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:jasi123@localhost:5432"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbName := "urlshortner"

	var exists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		fmt.Println(err)
	}

	if !exists {
		err = db.Exec("CREATE DATABASE urlshortner").Error
		if err != nil {
			log.Fatal(err)
		}
		log.Println("created database " + dbName)
	}

	db, err = gorm.Open(postgres.Open("postgres://postgres:jasi123@localhost:5432"+"/"+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&URls{})

	return db
}
