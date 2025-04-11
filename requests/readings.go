package requests

import (
	"UrbanWindComp/database"
	"fmt"
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

func GetLastReadings(c *fiber.Ctx) error {
	db := database.DBConn
	var readings []Reading
	db.Order("Epoch DESC").Limit(10).Find(&readings)
	return c.JSON(readings)
}

func GetReadings(c *fiber.Ctx) error {
	db := database.DBConn
	var readings []Reading
	db.Find(&readings)
	return c.JSON(readings)
}

func NewReading(c *fiber.Ctx) error {
	db := database.DBConn
	reading := new(Reading)
	fmt.Println(c.Body())
	if err := c.BodyParser(reading); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&reading)
	return c.JSON(reading)
}
