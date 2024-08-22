package usecase

import (
	"backend/internal/entities"
	"backend/internal/repositories"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserUsecase struct {
	DB             *gorm.DB
	UserRepository *repositories.UserRepository
}

func NewUserUsecase(db *gorm.DB, userRepository *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		DB:             db,
		UserRepository: userRepository,
	}
}

func (u *UserUsecase) Create() error {
	return u.UserRepository.Create(
		&entities.User{
			Name:      "test",
			ID:        uuid.NewString(),
			Email:     "gKkQz@example.com",
			Password:  "test",
			CreatedAt: time.Now(),
		},
	)
}
