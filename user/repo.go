package user

import (
	"errors"
	"log"

	"github.com/howkyle/uman"
	"gorm.io/gorm"
)

var CreationError = errors.New("unable to create new user")
var InvalidUser = errors.New("invalid user type passed")

type repository struct {
	DB *gorm.DB
}

func (r repository) Create(u uman.User) (interface{}, error) {

	user, ok := u.(User)
	if !ok {
		return nil, InvalidUser
	}

	res := r.DB.Create(&u)
	if res.Error != nil {
		log.Println(res.Error)
		return nil, CreationError
	}
	return user.ID, nil
}

func (r repository) Retrieve(id interface{}) (uman.User, error) {
	//can use id or struct for gorm
	return nil, nil
}

func (r repository) Delete(id interface{}) error {
	return nil
}

func NewRepository(db *gorm.DB) uman.UserRepository {
	return repository{DB: db}
}
