package main

import (
	"boilerplate/database"
	"boilerplate/handlers"
	"boilerplate/service"
	"os"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Connected with database
	database.Connect()
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = "5000"
	}

	// Create fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Statuscode defaults to 500
			code := fiber.StatusInternalServerError

			// Retreive the custom statuscode if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).JSON(fiber.Map{
				"message": err.Error(),
			})
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(500).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}

			// Return from handler
			return nil
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

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
	app.Get("/panic", handlers.PanicCheck)
	log.Fatal(app.Listen(":" + portEnv))
}
