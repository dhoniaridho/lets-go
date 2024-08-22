package entities

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"not null;column:name;type:varchar(255)"`
	Email     string    `gorm:"not null;column:email;type:varchar(255)"`
	Password  string    `gorm:"not null;column:password;type:varchar(255)"`
	CreatedAt time.Time `gorm:"not null;column:createdAt;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamp"`
}

func (u *User) TableName() string {
	return "users"
}
