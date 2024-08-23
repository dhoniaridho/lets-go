package route

import (
	"backend/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	AuthController *http.AuthController
}

func (r *RouteConfig) Setup() {
	r.SetupAuthRoute()
	r.SetupUserRoute()
}

func (r *RouteConfig) SetupAuthRoute() {
	r.App.Post("/v1/auth/sign-in", r.AuthController.SignIn)
}

func (r *RouteConfig) SetupUserRoute() {
	r.App.Post("/v1/user/create", r.UserController.Create)
}
