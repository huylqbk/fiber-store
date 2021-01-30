package main

import (
	"boilerplate/database"
	"boilerplate/handlers"
	"boilerplate/service"
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

	s := service.NewService("play.minio.io", "Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", false)

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.UserList)
	v1.Post("/users", handlers.UserCreate)
	v1.Post("/file", handlers.PushFile(s))
	app.Get("/download/:path", handlers.GetFile(s))

	// Setup static files
	app.Get("/", handlers.HealthCheck)
	log.Fatal(app.Listen(":" + portEnv))
}
