package event

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo Repository
}

//create a new event
func (s service) CreateEvent(c CreateEvent) (interface{}, error) {
	//validate struct then create

	id, err := s.repo.Create(c.Event())
	if err != nil {
		return 0, fmt.Errorf("failed to create event: %w", err)
	}
	return id, nil
}

func (s service) AddInvite(i Invitation) error {
	event, err := s.repo.Retrieve(i.EventID)
	if err != nil {
		return fmt.Errorf("failed to retrieve event: %w", err)
	}
	if (event.Invitations[i.UserID.Hex()] != Invitation{}) {
		return fmt.Errorf("user already invited")
	}
	event.Invitations[i.UserID.Hex()] = i
	err = s.repo.Update(event)
	if err != nil {
		return fmt.Errorf("failed to add event: %w", err)
	}
	return nil
}

func NewService(r Repository) Service {
	return service{repo: r}
}

//useCases

//used when creating a new event
type CreateEvent struct {
	UserID       primitive.ObjectID
	Title        string
	Latitude     int
	Longitude    int
	LocationName string
	Time         time.Time
}

//maps usecase to an event
func (c CreateEvent) Event() Event {
	return Event{UserID: c.UserID, Title: c.Title, Location: Location{Latitude: c.Latitude, Longitude: c.Longitude, Name: c.LocationName}, Invitations: make(map[string]Invitation)}
}

type CreateInvitation struct {
	UserID   primitive.ObjectID
	EventID  primitive.ObjectID
	Accepted bool
}

func (c CreateInvitation) Invitation() Invitation {
	return Invitation{UserID: c.UserID, EventID: c.EventID, Accepted: false}
}
