package repositories

import (
	"backend/internal/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entities.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository[entities.User]{
			DB: db,
		},
	}
}
