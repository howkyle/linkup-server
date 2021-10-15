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
	return Event{UserID: c.UserID, Title: c.Title, Location: Location{Latitude: c.Latitude, Longitude: c.Longitude, Name: c.LocationName}}
}
