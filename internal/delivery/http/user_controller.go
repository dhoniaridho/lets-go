package http

import (
	"backend/internal/entities"
	"backend/internal/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	UseCase *usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{
		UseCase: usecase,
	}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	err := c.UseCase.UserRepository.Create(
		&entities.User{
			Name:      "test",
			Email:     "gKkQz@example.com",
			Password:  "test",
			CreatedAt: time.Now(),
			ID:        uuid.NewString(),
		},
	)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendString("Hello, World!")
}
