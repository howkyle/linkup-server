package event

import "go.mongodb.org/mongo-driver/mongo"

type mongorepo struct {
	db *mongo.Database
}

func (r mongorepo) Create(e Event) (interface{}, error) {
	return 0, nil
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
