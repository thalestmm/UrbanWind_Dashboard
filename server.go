package main

import (
	"UrbanWindComp/database"
	"UrbanWindComp/requests"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New())

	initDatabase()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func initDatabase() {
	var err error
	db := database.DBConn

	// Open connection to DB
	db, err = gorm.Open(sqlite.Open("readings.db"))
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate tables
	err1 := db.AutoMigrate(&requests.Reading{})
	if err1 != nil {
		panic(err1)
	}
}
