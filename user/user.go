package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Service interface {
	Register(u User) (uint, error)
	Login() error
	Logout() error
}

type Repo interface {
	Add(u User) (uint, error)
	Get(id uint) (User, error)
	Delete(id uint) error
}
