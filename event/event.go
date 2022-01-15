//handles events
package event

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Location
	Time time.Time `bson:"time,omitempty"`
}

type Location struct {
	Latitude     int    `bson:"latitude,omitempty"`
	Longitude    int    `bson:"longitude,omitempty"`
	LocationName string `bson:"location_name,omitempty"`
}

type Service interface {
	CreateEvent(e Event) (interface{}, error)
	Event(id interface{}) (Event, error)
}

type Repository interface {
	//adds a new event record to the database
	Create(e Event) (interface{}, error)
	//retrieves an event record from the database
	Retrieve(e interface{}) (Event, error)
	//updates an existing event record
	Update(e Event) error
	//deletes an existing event record
	Delete(e Event) error
}
