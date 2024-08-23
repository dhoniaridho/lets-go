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
	r, err := c.UseCase.SignIn()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "success",
			"data":    r,
		},
	)
}
