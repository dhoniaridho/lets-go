package usecase

import (
	"backend/internal/entities"
	"backend/internal/repositories"
	"backend/lib/jwt"
	"errors"

	"gorm.io/gorm"
)

type AuthUsecase struct {
	DB             *gorm.DB
	UserRepository *repositories.UserRepository
}

func NewAuthUseCase(db *gorm.DB, UserRepository *repositories.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		DB:             db,
		UserRepository: UserRepository,
	}
}

func (u *AuthUsecase) SignIn() (string, error) {
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
