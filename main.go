package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, world!")

	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello world!",
		})
	})

	port := 6000
	addr := fmt.Sprintf(":%d", port)

	err := app.Listen(addr)
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}
