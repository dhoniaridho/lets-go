package middleware

import (
	"backend/internal/models"
	"backend/internal/utils"
	"backend/lib/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &models.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}

		if request.Token == "NOT_FOUND" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&utils.Response{
				Status:  fiber.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			})
		}

		data, err := jwt.Verify(strings.Split(request.Token, " ")[1])

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&utils.Response{
				Status:  fiber.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			})
		}

		ctx.Locals("auth", data)

		return ctx.Next()
	}
}
