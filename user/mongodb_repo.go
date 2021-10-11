package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var CreationError = errors.New("unable to add user to repository")
var InvalidUser = errors.New("invalid user type passed")
var NotFound = errors.New("user not found")

const Collection = "users"

type mongorepo struct {
	db *mongo.Database
}

func (r mongorepo) Create(u User) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := r.db.Collection(Collection).InsertOne(ctx, u)
	if err != nil {
		log.Println(err)
		return nil, CreationError
	}
	return res.InsertedID, nil
}

func (r mongorepo) Retrieve(u User) (User, error) {
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	err := r.db.Collection(Collection).FindOne(ctx, u).Decode(&user)
	if err != nil {
		return User{}, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}

func (r mongorepo) Delete(id interface{}) error {
	return nil
}

func NewMongoRepository(db *mongo.Database) Repository {
	return mongorepo{db: db}
}
