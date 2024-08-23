package config

import (
	"backend/internal/delivery/http"
	"backend/internal/delivery/http/route"
	"backend/internal/repositories"
	"backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App *fiber.App
	DB  *gorm.DB
}

func Bootstrap(config *BootstrapConfig) {

	userRepository := repositories.NewUserRepository(config.DB)

	// setup use cases
	userUseCase := usecase.NewUserUsecase(config.DB, userRepository)
	authUseCase := usecase.NewAuthUseCase(config.DB, userRepository)

	// setup controllers
	userController := http.NewUserController(userUseCase)
	authController := http.NewAuthController(authUseCase)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		AuthController: authController,
	}

	routeConfig.Setup()
}
