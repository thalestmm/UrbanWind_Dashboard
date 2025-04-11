package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DBConn *gorm.DB

func initDatabase() {
	var err error

	dbPath := os.Getenv("DB_PATH")

	DBConn, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{}) // or postgres, etc.
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("✅ Database connected")

	DBConn.AutoMigrate(&Reading{}) // make sure Reading is defined
	log.Println("✅ Database migrated")

}
