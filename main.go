package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main(){
	// Connect to the database
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("Database connection established:", db)
	// Create a new Fiber app
	app := fiber.New()

	// test endpoint
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}