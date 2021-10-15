package invitation

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo Repository
}

func (s service) AddInvitation(i Invitation) (interface{}, error) {
	invitation, err := s.repo.Retrieve(Invitation{EventID: i.EventID, UserID: i.UserID})
	if (invitation != Invitation{}) {
		return nil, fmt.Errorf("duplicate invitation")
	}
	id, err := s.repo.Create(i)
	if err != nil {
		return nil, fmt.Errorf("failed to add invitation: %w", err)
	}
	return id, nil
}

func (s service) Invitation(id interface{}) (Invitation, error) {

	invitationID, ok := id.(primitive.ObjectID)
	if !ok {
		return Invitation{}, fmt.Errorf("invalid invitation id")
	}
	i, err := s.repo.Retrieve(Invitation{ID: invitationID})
	if err != nil {
		return Invitation{}, err
	}
	return i, nil
}

func (s service) InvitationsByEvent(eid interface{}) ([]Invitation, error) {
	eventid, ok := eid.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid event id")
	}
	i, err := s.repo.RetrieveMultiple(Invitation{EventID: eventid})
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (s service) InvitationsByUser(uid interface{}) ([]Invitation, error) {
	userid, ok := uid.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid user id")
	}
	i, err := s.repo.RetrieveMultiple(Invitation{UserID: userid})
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (s service) AcceptInvitation(id interface{}) error {
	invitationID, ok := id.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("invalid invitation id")
	}
	i, err := s.repo.Retrieve(Invitation{ID: invitationID})
	if err != nil {
		return fmt.Errorf("failed to accept invitation: %v", err)
	}
	i.Accepted = true
	err = s.repo.Update(i)
	if err != nil {
		return fmt.Errorf("failed to update invitation: %v", err)
	}
	return nil
}

func NewService(r Repository) Service {
	return service{repo: r}
}

type CreateInvitation struct {
	UserID   primitive.ObjectID
	EventID  primitive.ObjectID
	Accepted bool
}

func (c CreateInvitation) Invitation() Invitation {
	return Invitation{UserID: c.UserID, EventID: c.EventID, Accepted: false}
}
