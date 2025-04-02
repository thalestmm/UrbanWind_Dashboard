package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "UrbanWind",
		})
	})
	app.Get("/update-chart", func(c *fiber.Ctx) error {
		return c.Render("chart", fiber.Map{
			// TODO: Get last 10 readings from DB
			"data": nil,
		})
	})
	app.Post("/reading", func(c *fiber.Ctx) error {
		r := new(ReadingRequest)
		if err := c.BodyParser(r); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		// TODO: Add reading to DB
		return nil
	})
}
