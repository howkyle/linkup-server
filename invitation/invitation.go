package invitation

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invitation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty"`
	EventID  primitive.ObjectID `bson:"event_id,omitempty"`
	Summary  string             `bson:"summary,omitempty"`
	Accepted bool               `bson:"accepted,omitempty"`
}

type Service interface {
	AddInvitation(i Invitation) (interface{}, error)
	Invitation(filter interface{}) (Invitation, error)
	InvitationsByUser(userid interface{}) ([]Invitation, error)
	InvitationsByEvent(eventid interface{}) ([]Invitation, error)
	AcceptInvitation(userid, id interface{}) error
}

type Repository interface {
	//adds a new invitation document to the database
	Create(i Invitation) (interface{}, error)
	//retrieves an invitation document from the database
	Retrieve(i interface{}) (Invitation, error)
	//retrieves a slice of events
	RetrieveMultiple(filter interface{}) ([]Invitation, error)
	//updates an existing invitation document
	Update(i Invitation) error
	//deletes an existing invitation document
	Delete(i Invitation) error
}
