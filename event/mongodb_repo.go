package event

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (r mongorepo) Retrieve(id interface{}) (Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var event Event
	log.Printf("%v", id)
	filter := bson.D{{Key: "_id", Value: id}}
	err := r.db.Collection(collection).FindOne(ctx, filter).Decode(&event)
	if err != nil {
		return Event{}, fmt.Errorf("failed to retrieve event: %w", err)
	}
	return event, nil
}

func (r mongorepo) Update(e Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: e.ID}}
	update := bson.D{{Key: "$set", Value: e}}
	res, err := r.db.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}
	if res.ModifiedCount != 1 {
		return fmt.Errorf("event document not modified")
	}
	return nil
}

func (r mongorepo) Delete(e Event) error {
	return nil
}

func NewMongoRepository(db *mongo.Database) Repository {
	return mongorepo{db: db}
}
