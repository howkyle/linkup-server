package user

import (
	"time"

	"github.com/howkyle/linkup-server/event"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Events    []event.Event
}

type Service interface {
	Register(u User) (uint, error)
	Login(u User) (interface{}, error)
	Logout() error
}

type Repository interface {
	Delete(id interface{}) error
	Retrieve(u User) (User, error)
	Create(u User) (interface{}, error)
}
