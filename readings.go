package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ReadingRequest struct {
	Value float64 `json:"value" form:"value" xml:"value"`
	Epoch int64   `json:"epoch" form:"epoch" xml:"epoch"`
}

type Reading struct {
	gorm.Model
	Value float64 `json:"value"`
	Epoch int64   `json:"epoch"`
}

func GetLastReadings(c *fiber.Ctx) ([]Reading, error) {
	var readings []Reading
	result := DBConn.Order("Epoch DESC").Limit(10).Find(&readings)
	return readings, result.Error
}

func NewReading(c *fiber.Ctx) error {
	db := DBConn

	var readingRequest ReadingRequest

	// Parse the JSON body into ReadingRequest
	if err := c.BodyParser(&readingRequest); err != nil {
		return c.Status(400).SendString("Invalid input: " + err.Error())
	}

	// Map ReadingRequest to Reading
	dbReading := Reading{
		Value: readingRequest.Value,
		Epoch: readingRequest.Epoch,
	}

	// Save to DB
	if err := db.Create(&dbReading).Error; err != nil {
		return c.Status(500).SendString("Database error: " + err.Error())
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"id":      dbReading.ID,
		"value":   dbReading.Value,
		"epoch":   dbReading.Epoch,
	})
}
