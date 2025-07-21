package controller

import (

	"example.com/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func ProductsController(app fiber.Router, db *gorm.DB){
	// Define the product routes
	app.Get("/", func(c *fiber.Ctx) error {
		var products []model.Product
		if err := db.Find(&products).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching products")
		}
		return c.JSON(products)
	})

}