package user

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

var CreationError = errors.New("unable to create new user")
var InvalidUser = errors.New("invalid user type passed")
var NotFound = errors.New("user not found")

type repository struct {
	DB *gorm.DB
}

func (r repository) Create(u User) (interface{}, error) {

	res := r.DB.Create(&u)
	if res.Error != nil {
		log.Println(res.Error)
		return nil, CreationError
	}
	return u.ID, nil
}

func (r repository) Retrieve(u User) (User, error) {
	//can use id or struct for gorm
	var user User
	res := r.DB.Where(&u).First(&user)
	if res.Error != nil {
		log.Println(res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return User{}, NotFound
		}
		return User{}, res.Error
	}
	return user, nil
}

func (r repository) Delete(id interface{}) error {
	return nil
}

func NewRepository(db *gorm.DB) Repository {
	return repository{DB: db}
}
