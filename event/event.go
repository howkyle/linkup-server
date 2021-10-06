//handles events
package event

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UserID      uint
	Title       string
	Location    Location
	Time        time.Time
	Invitations map[uint]Invitation
}

func (e *Event) AddInvitation(i Invitation) error {
	if (e.Invitations[i.UserID] != Invitation{}) {
		return errors.New("user already invited")
	}
	e.Invitations[i.UserID] = i
	return nil
}

type Location struct {
	Latitude  int
	Longitude int
	Name      string
}

type Repository interface {
	//adds a new event record to the database
	Create(e Event) (uint, error)
	//retrieves an event record from the database
	Retrieve(e interface{}) (Event, error)
	//updates an existing event record
	Update(e Event) error
	//deletes an existing event record
	Delete(e interface{}) error
}
