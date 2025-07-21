package main

import (
	"github.com/gofiber/fiber/v2"
)

func main(){
	// Create a new Fiber app
	app := fiber.New()

	// test endpoint
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}