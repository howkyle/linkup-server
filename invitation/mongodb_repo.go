package invitation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "invitations"

var CreationError = errors.New("failed to create invitation")

type mongorepo struct {
	db *mongo.Database
}

func (r mongorepo) Create(i Invitation) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := r.db.Collection(collection).InsertOne(ctx, i)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", CreationError, err)
	}
	return res.InsertedID, nil
}

func (r mongorepo) Retrieve(id interface{}) (Invitation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var i Invitation
	filter := bson.D{{Key: "_id", Value: id}}
	err := r.db.Collection(collection).FindOne(ctx, filter).Decode(&i)
	if err != nil {
		return Invitation{}, fmt.Errorf("failed to retrieve invitation: %w", err)
	}
	return i, nil
}

func (r mongorepo) RetrieveMultiple(filter interface{}) ([]Invitation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var invitations []Invitation = make([]Invitation, 0)
	cur, err := r.db.Collection(collection).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve invitations: %v", err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var i Invitation
		err = cur.Decode(&i)
		if err != nil {
			return nil, fmt.Errorf("failed to decode results: %v", err)
		}
		invitations = append(invitations, i)
	}
	err = cur.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve invitations: %w", err)
	}
	return invitations, nil
}

func (r mongorepo) Update(i Invitation) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: i.ID}}
	update := bson.D{{Key: "$set", Value: i}}
	res, err := r.db.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update invitation document: %w", err)
	}
	if res.ModifiedCount != 1 {
		return fmt.Errorf("invitation document not modified")
	}
	return nil
}

func (r mongorepo) Delete(e Invitation) error {
	return nil
}

func NewMongoRepository(db *mongo.Database) Repository {
	return mongorepo{db: db}
}
