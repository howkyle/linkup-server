package user

import (
	"errors"
	"log"

	"github.com/howkyle/uman"
)

type UserSignup struct {
	Username string
	Email    string
	Password string
}

func (u UserSignup) User() User {
	return User{Username: u.Username, Email: u.Email, Password: u.Password}
}

var RegistrationError = errors.New("user registration failed")
var InvalidID = errors.New("invalid user id type")

type service struct {
	repo        uman.UserRepository
	userManager uman.UserManager
	authManager uman.AuthManager
}

func (s service) Register(u User) (uint, error) {
	id, err := s.userManager.Create(u)
	if err != nil {
		log.Println(err)
		return 0, RegistrationError
	}
	uid, ok := id.(uint)
	if !ok {
		return 0, InvalidID
	}

	return uid, nil
}

func (s service) Login(u User, p string) error {
	s.userManager.Retrieve(u.Username)

	return nil
}

func (s service) Logout() error {
	// s.userManager.Retrieve(u.Username)

	return nil
}

func NewService(r uman.UserRepository, a uman.AuthManager, u uman.UserManager) Service {
	return service{repo: r, authManager: a, userManager: u}
}
