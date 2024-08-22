package http

import (
	"backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	UseCase *usecase.AuthUsecase
}

func NewAuthController(usecase *usecase.AuthUsecase) *AuthController {
	return &AuthController{
		UseCase: usecase,
	}
}

func (c *AuthController) SignIn(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
