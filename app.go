package main

import (
	"boilerplate/database"
	"boilerplate/handlers"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	port = flag.String("port", ":5000", "Port to listen on")
)

func main() {
	// Connected with database
	database.Connect()

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
	log.Fatal(app.Listen(*port))
}
