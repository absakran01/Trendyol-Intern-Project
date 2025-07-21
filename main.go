package main

import (
	"log"

	"example.com/controller"
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

	//product endpoints
	controller.ProductsController(app.Group("/api/product"), db)


	log.Println("Listening on port 3000...")
	app.Listen(":3000")
}  