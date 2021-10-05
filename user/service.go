package user

import (
	"errors"
	"fmt"
	"log"

	"github.com/howkyle/authman"
)

var RegistrationError = errors.New("user registration failed")
var InvalidID = errors.New("invalid user id type")
var LoginError = errors.New("login failure")

type UserSignup struct {
	Username string
	Email    string
	Password string
}

func (u UserSignup) User() User {
	return User{Username: u.Username, Email: u.Email, Password: u.Password}
}

type UserLogin struct {
	Username string
	Password string
}

func (ul UserLogin) User() User {
	return User{Username: ul.Username, Password: ul.Password}
}

type service struct {
	repo Repository
	// userManager uman.UserManager
	authManager authman.AuthManager
}

func (s service) Register(u User) (uint, error) {
	hashedPass, err := authman.NewUserPassCredentials(u.Username, u.Password).Hash()
	if err != nil {
		log.Println(err)
		return 0, RegistrationError
	}
	u.Password = hashedPass
	id, err := s.repo.Create(u)
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

func (s service) Login(u User) (interface{}, error) {
	user, err := s.repo.Retrieve(User{Username: u.Username})
	if err != nil {
		return nil, fmt.Errorf("%v: %w", LoginError, err)
	}
	c := authman.NewUserPassCredentials(fmt.Sprintf("%v", user.ID), user.Password)
	auth, err := s.authManager.Authenticate(c, u.Password)

	if err != nil {
		return nil, fmt.Errorf("%v: %w", LoginError, err)
	}
	return auth.AsCookie(), nil
}

func (s service) Logout() error {
	// s.userManager.Retrieve(u.Username)

	return nil
}

func NewService(r Repository, a authman.AuthManager) Service {
	return service{repo: r, authManager: a}
}
