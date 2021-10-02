package user

import (
	"time"

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
}

func (u User) GetID() interface{} {
	return u.ID
}
func (u User) GetUsername() string {
	return u.Username
}
func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetEmail() string {
	return u.Email
}

type Service interface {
	Register(u User) (uint, error)
	Login(u User, p string) error
	Logout() error
}
