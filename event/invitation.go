package event

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invitation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty"`
	EventID  primitive.ObjectID `bson:"event_id,omitempty"`
	Accepted bool               `bson:"accepted,omitempty"`
}
