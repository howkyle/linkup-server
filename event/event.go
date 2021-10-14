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
	Time        time.Time             `bson:"time,omitempty"`
	Invitations map[string]Invitation `bson:"invitations, omitempty"`
}

type Location struct {
	Latitude  int    `bson:"latitute,omitempty"`
	Longitude int    `bson:"longitude,omitempty"`
	Name      string `bson:"name,omitempty"`
}

type Service interface {
	CreateEvent(e CreateEvent) (interface{}, error)
	AddInvite(i Invitation) error
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
