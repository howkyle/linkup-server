package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type Service interface {
	Register(u User) (interface{}, error)
	Login(u User) (interface{}, error)
	Logout() error
}

type Repository interface {
	Create(u User) (interface{}, error)
	Retrieve(u User) (User, error)
	Update(u User) error
	Delete(id interface{}) error
}
