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

		// get jwt claims
		claims, ok := data.Get("claims")
		if !ok {
			// handle the case where the key was not found
			return ctx.Status(fiber.StatusUnauthorized).JSON(&utils.Response{
				Status:  fiber.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			})
		}

		claimsMap := claims.(map[string]interface{})

		ctx.Locals("auth", &jwt.Claims{
			UserID:   claimsMap["userId"].(string),
			Username: claimsMap["username"].(string),
		})
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *jwt.Claims {
	return ctx.Locals("auth").(*jwt.Claims)
}
