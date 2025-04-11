package main

import (
	"UrbanWindComp/requests"
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
		payload := requests.GetLastReadings(c)
		fmt.Println(payload)

		return c.Render("chart", fiber.Map{
			// TODO: Get last 10 readings from DB
			"data": nil,
		})
	})

	//app.Post("/reading", func(c *fiber.Ctx) error {
	//	r := new(requests.ReadingRequest)
	//	if err1 := c.BodyParser(r); err1 != nil {
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": err1,
	//		})
	//	}
	//	err2 := requests.NewReading(c)
	//	if err2 != nil {
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": err2,
	//		})
	//	}
	//	return nil
	//})

	app.Post("/reading", requests.NewReading)
}
