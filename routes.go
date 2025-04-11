package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "UrbanWind",
		})
	})

	app.Get("/update-chart", func(c *fiber.Ctx) error {
		payload, err := GetLastReadings(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		fmt.Println(payload)

		return c.Render("chart", fiber.Map{
			// TODO: Get last 10 readings from DB
			"Data": nil,
		})
	})

	app.Post("/reading", NewReading)

}
