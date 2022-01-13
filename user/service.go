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

type service struct {
	repo        Repository
	authManager authman.AuthManager
}

//takes and new user and hashes pass and stores to database
func (s service) Register(u User) (interface{}, error) {
	err := checkDuplicate(s.repo, u)
	if err != nil {
		return nil, err
	}
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

	return id, nil
}

//checks if a user with the given username or email exist in the database
func checkDuplicate(r Repository, u User) error {
	_, err := r.Retrieve(User{Username: u.Username, Email: u.Email})
	if err == nil {
		return fmt.Errorf("duplicate user")
	}
	return nil
}

//takes a user and authenticates the user and returns auth token/cookie
func (s service) Login(u User) (interface{}, error) {
	user, err := s.repo.Retrieve(User{Username: u.Username})
	if err != nil {
		return nil, fmt.Errorf("%v: %w", LoginError, err)
	}
	c := authman.NewUserPassCredentials(user.ID.Hex(), user.Password)
	auth, err := s.authManager.Authenticate(c, u.Password)

	if err != nil {
		return nil, fmt.Errorf("%v: %w", LoginError, err)
	}
	return auth.AsCookie(), nil
}

func (s service) Logout() error {
	return nil
}

func NewService(r Repository, a authman.AuthManager) Service {
	return service{repo: r, authManager: a}
}
