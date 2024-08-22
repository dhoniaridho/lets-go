package repositories

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(t *T) error {
	return r.DB.Create(t).Error
}
