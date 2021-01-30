package main

import (
	"boilerplate/database"
	"boilerplate/handlers"
	"os"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Connected with database
	database.Connect()
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = "5000"
	}

	// Create fiber app
	app := fiber.New(fiber.Config{})

	// Middleware
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.UserList)
	v1.Post("/users", handlers.UserCreate)

	// Setup static files
	app.Get("/", handlers.HealthCheck)
	log.Fatal(app.Listen(":" + portEnv))
}
