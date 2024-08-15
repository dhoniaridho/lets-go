package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App
}

func (r *RouteConfig) Setup() {

	r.App.Get("/", func(c *fiber.Ctx) error {
		body := c.Query("name")
		return c.SendString(fmt.Sprintf("Hello, %s!", body))
	})

}
