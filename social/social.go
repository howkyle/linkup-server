//handles social interactions of the the users
package social

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Friendship struct {
	ID          primitive.ObjectID
	Users       []primitive.ObjectID
	DateCreated time.Time
}

type FriendInvite struct {
	ID          primitive.ObjectID
	DateCreated time.Time
	Creator     string
	CreatorID   primitive.ObjectID
	RecipientID primitive.ObjectID
	Accepted    bool
	Rejected    bool
}

type Service interface {
	InviteFriend(userID, friendID interface{}) error
	Invitations(userID interface{}) ([]FriendInvite, error)
	AcceptInvite(userID, inviteID interface{}) error
	Friends(userID interface{}) ([]interface{}, error)
}

type Repository interface {
	InviteRepo
	FriendshipRepo
}

type InviteRepo interface {
	CreateInvite(f FriendInvite) (interface{}, error)
	Invite(id interface{}) (FriendInvite, error)
	InvitesByRecipient(id interface{}) ([]FriendInvite, error)
	DeleteInvite(id interface{}) error
}

type FriendshipRepo interface {
	CreateFriendship(f Friendship) (interface{}, error)
	Friendship(userA, userB interface{}) (Friendship, error)
	Friendships(userId interface{}) ([]Friendship, error)
	DeleteFriendship(id interface{}) error
}
