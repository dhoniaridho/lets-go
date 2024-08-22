package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgresql@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
