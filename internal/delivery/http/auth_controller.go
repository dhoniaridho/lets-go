package http

import (
	"backend/internal/delivery/http/middleware"
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
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusUnauthorized,
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&utils.Response{
			Data:    r,
			Message: "success",
			Status:  fiber.StatusOK,
		},
	)
}

func (c *AuthController) SignUp(ctx *fiber.Ctx) error {
	request := new(models.SignUpRequest)
	ctx.BodyParser(request)
	res, err := utils.ValidateBody(c.UseCase.Validate, request)

	if err != nil {
		println("Validation error: ", err)
		return ctx.Status(res.Status).JSON(res)
	}

	data, err := c.UseCase.SignUp(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.Response{
				Data:    nil,
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
				Errors:  []utils.ErrorResponse{},
			},
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.BuildResponse(
			&utils.Response{
				Data:    data,
				Message: "success",
				Status:  fiber.StatusOK,
			},
		),
	)
}

func (c *AuthController) GetProfile(ctx *fiber.Ctx) error {

	auth := middleware.GetUser(ctx)

	user, err := c.UseCase.GetProfile(auth.UserID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.Response{
				Data:    nil,
				Message: err.Error(),
				Status:  fiber.StatusBadRequest,
				Errors:  []utils.ErrorResponse{},
			},
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.Response{
			Data:    user,
			Message: "success",
			Status:  fiber.StatusOK,
		},
	)
}
