package event

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "events"

var CreationError = errors.New("failed to add event to database")

type mongorepo struct {
	db *mongo.Database
}

func (r mongorepo) Create(e Event) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := r.db.Collection(collection).InsertOne(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", CreationError, err)
	}
	return res.InsertedID, nil
}

func (r mongorepo) Retrieve(e interface{}) (Event, error) {
	return Event{}, nil
}

func (r mongorepo) Update(e Event) error {
	return nil
}

func (r mongorepo) Delete(e interface{}) error {
	return nil
}

func NewMongoRepository(db *mongo.Database) Repository {
	return mongorepo{db: db}
}
