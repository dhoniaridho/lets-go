package route

import (
	"backend/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	AuthController *http.AuthController
	AuthMiddlewre  fiber.Handler
}

func (r *RouteConfig) Setup() {
	r.SetupAuthRoute()
	r.SetupUserRoute()
}

func (r *RouteConfig) SetupAuthRoute() {
	r.App.Post("/v1/auth/sign-up", r.AuthController.SignUp)
	r.App.Post("/v1/auth/sign-in", r.AuthController.SignIn)
}

func (r *RouteConfig) SetupUserRoute() {
	r.App.Use(r.AuthMiddlewre)
	r.App.Post("/v1/user/create", r.UserController.Create)
	r.App.Get("/v1/auth/profile", r.AuthController.GetProfile)

}
