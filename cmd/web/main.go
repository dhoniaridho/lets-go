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

	app := fiber.New(
		fiber.Config{
			// Prefork:           true,
			EnablePrintRoutes: true,
		},
	)

	db := config.NewDatabase()
	v := config.NewValidator()

	config.Bootstrap(
		&config.BootstrapConfig{
			App:       app,
			DB:        db,
			Validator: v,
		},
	)

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
