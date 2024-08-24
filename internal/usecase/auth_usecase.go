package usecase

import (
	"backend/internal/entities"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/lib/jwt"
	"errors"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthUsecase struct {
	DB             *gorm.DB
	UserRepository *repositories.UserRepository
	Validate       *validator.Validate
}

func NewAuthUseCase(db *gorm.DB, UserRepository *repositories.UserRepository, Validate *validator.Validate) *AuthUsecase {
	return &AuthUsecase{
		DB:             db,
		UserRepository: UserRepository,
		Validate:       Validate,
	}
}

func (u *AuthUsecase) SignIn(data *models.SignInRequest) (string, error) {
	var user entities.User

	err := u.UserRepository.First(&user)

	if err != nil {
		return "", errors.New("user not found")
	}

	token, err := jwt.Sign()

	if err != nil {
		return "", err
	}

	return token, nil

}
