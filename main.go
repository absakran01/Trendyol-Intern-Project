package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"example.com/controller"
	"example.com/model"
	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to the database
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("Database connection established:", db.Name())

	// kafka topic
	topic := "postgres.public.products"

	// create consumer then start
	broker := []string{"127.0.0.1:9092"}
	consumer, err := connectToConsumer(broker)
	if err != nil {
		panic(err)
	}

	partConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Error starting partition consumer:", err)
	}
	log.Printf("consumer is running with topic: %s\n", topic)

	// handle os signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// make go routines
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-partConsumer.Errors():
				log.Println(err)
			case msg := <-partConsumer.Messages():
				{
					product := toProduct(msg.Value)
					if product.Title == "" {
						log.Println("Received empty product title")
					} else {
						log.Printf("update detected for product: %s\n", product.Title)
					}
				}
			case <-sigChan:
				{
					log.Println("interrupt detected")
					doneCh <- struct{}{}
				}
			}
		}
	}()
	<-doneCh
	log.Println("Shutting down consumer...")

	// close consumers on exit
	err = consumer.Close()
	if err != nil {
		panic(err)
	}

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

func toProduct(value []byte) model.Product {
	var envelope struct {
		Payload struct {
			After json.RawMessage `json:"after"`
		} `json:"payload"`
	}
	if err := json.Unmarshal(value, &envelope); err != nil {
		log.Println("Error unmarshalling envelope:", err)
		return model.Product{}
	}
	var product model.Product
	if len(envelope.Payload.After) == 0 || string(envelope.Payload.After) == "null" {
		log.Println("No 'after' data in payload")
		return model.Product{}
	}
	if err := json.Unmarshal(envelope.Payload.After, &product); err != nil {
		log.Println("Error unmarshalling product:", err)
		return model.Product{}
	}
	return product
}
