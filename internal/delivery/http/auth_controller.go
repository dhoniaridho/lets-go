package http

import (
	"backend/internal/models"
	"backend/internal/usecase"
	"backend/internal/utils"

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

	request := new(models.SignInRequest)
	ctx.BodyParser(request)

	res, err := utils.ValidateBody(c.UseCase.Validate, request)

	if err != nil {
		println("Validation error: ", err)
		return ctx.Status(res.Status).JSON(res)
	}

	r, err := c.UseCase.SignIn(request)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.BuildResponse(
			&utils.Response{
				Data: map[string]interface{}{
					"token": r,
				},
				Message: "success",
				Status:  fiber.StatusOK,
			},
		),
	)
}
