package user

import (
	"github.com/howkyle/linkup-server/event"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Events   []event.Event
}

type Service interface {
	Register(u User) (interface{}, error)
	Login(u User) (interface{}, error)
	Logout() error
}

type Repository interface {
	Delete(id interface{}) error
	Retrieve(u User) (User, error)
	Create(u User) (interface{}, error)
}
