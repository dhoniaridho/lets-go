package main

import (
	_ "backend/docs"
	"backend/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()

	db := config.NewDatabase()

	config.Bootstrap(
		&config.BootstrapConfig{
			App: app,
			DB:  db,
		},
	)

	// Log all mapped routes
	for _, route := range app.GetRoutes() {
		log.Printf("Route mapped: %s %s\n", route.Method, route.Path)
	}

	// setup cors

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		},
	))

	app.Listen(":2000")
}
