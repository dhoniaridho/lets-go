package repositories

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(t *T) error {
	return r.DB.Create(t).Error
}

func (r *Repository[T]) Update(t *T) error {
	return r.DB.Updates(t).Error
}

func (r *Repository[T]) Delete(t *T) error {
	return r.DB.Delete(t).Error
}

func (r *Repository[T]) First(t *T) error {
	return r.DB.First(t).Error
}

func (r *Repository[T]) Find(t *T) error {
	return r.DB.Find(t).Error
}
