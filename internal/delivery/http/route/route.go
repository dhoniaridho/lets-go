package route

import (
	"backend/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
}

func (r *RouteConfig) Setup() {
	r.SetupAuthRoute()
}

func (r *RouteConfig) SetupAuthRoute() {
	r.App.Post("/v1/user/create", r.UserController.Create)
}
