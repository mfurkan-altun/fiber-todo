package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mfurkan-altun/fiber-todo/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// setup routes
	setupRoutes(app) // new

	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "You are at the endpoint 😉 (localhost:8000)",
		})
	})

	// Listen on server 8000 and catch error if any
	err := app.Listen(":8000")

	// handle error
	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/furk", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "You are at the endpoint 😉 (localhost:8000/furk)",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉 (localhost:8000/api)",
		})
	})

	// connect todo routes
	routes.TodoRoute(api.Group("/todos"))
}