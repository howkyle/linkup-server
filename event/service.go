package event

import (
	"fmt"
)

type service struct {
	repo Repository
}

//create a new event
func (s service) CreateEvent(c Event) (interface{}, error) {

	// check if dupilcate exists
	id, err := s.repo.Create(c)
	if err != nil {
		return 0, fmt.Errorf("failed to create event: %w", err)
	}
	return id, nil
}

//retrieves an event using the event id
func (s service) Event(id interface{}) (Event, error) {

	e, err := s.repo.Retrieve(id)
	if err != nil {
		return Event{}, fmt.Errorf("failed to retrieve event: %v", err)
	}
	return e, nil
}

func NewService(r Repository) Service {
	return service{repo: r}
}
