package usecase

import "gorm.io/gorm"

type UserUsecase struct {
	DB *gorm.DB
}

func NewUserUsecase(db *gorm.DB) *UserUsecase {
	return &UserUsecase{
		DB: db,
	}
}
