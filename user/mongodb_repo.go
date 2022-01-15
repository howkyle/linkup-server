package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var CreationError = errors.New("unable to add user to repository")
var InvalidUser = errors.New("invalid user type passed")
var NotFound = errors.New("user not found")

const Collection = "users"

type update struct {
	set User `bson:"$set"`
}
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

//Retrieves a user from the database matching the username or email address of the passed user
func (r mongorepo) Retrieve(u User) (User, error) {
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	filter := bson.D{{
		Key: "$or",
		Value: bson.A{
			bson.D{{Key: "username", Value: u.Username}},
			bson.D{{Key: "email", Value: u.Email}},
		},
	}}
	defer cancel()
	err := r.db.Collection(Collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}

func (r mongorepo) Update(u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: u.ID}}
	update := bson.D{{Key: "$set", Value: u}} //todo experiment with struct

	res, err := r.db.Collection(Collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	if res.ModifiedCount != 1 {
		return fmt.Errorf("failed to update user: no record modified")
	}

	return nil
}

func (r mongorepo) Delete(id interface{}) error {
	return nil
}

func NewMongoRepository(db *mongo.Database) Repository {
	return mongorepo{db: db}
}
