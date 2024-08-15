package main

import (
	"backend/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	config.Bootstrap(
		&config.BootstrapConfig{
			App: app,
		},
	)

	// Log all mapped routes
	for _, route := range app.GetRoutes() {
		log.Printf("Route mapped: %s %s\n", route.Method, route.Path)
	}

	app.Listen(":3000")
}
