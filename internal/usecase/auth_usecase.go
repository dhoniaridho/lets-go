package usecase

import "gorm.io/gorm"

type AuthUsecase struct {
	DB *gorm.DB
}

func NewAuthUseCase(db *gorm.DB) *AuthUsecase {
	return &AuthUsecase{
		DB: db,
	}
}
