package social

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const InviteCollection = "friend_invites"
const FriendshipCollection = "friendships"

type repo struct {
	db *mongo.Database
}

//adds a new FriendInvite document to the collection
func (r repo) CreateInvite(f FriendInvite) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := r.db.Collection(InviteCollection).InsertOne(ctx, f)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

//retrieves a Friendship Invite document from the collection using an id
func (r repo) Invite(id interface{}) (FriendInvite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var f FriendInvite
	filter := bson.D{{Key: "_id", Value: id}}
	err := r.db.Collection(InviteCollection).FindOne(ctx, filter).Decode(&f)
	if err != nil {
		return FriendInvite{}, err
	}
	return f, nil
}

//retrieves a slice of Friendship Invites from the collection using a recipient id
func (r repo) InvitesByRecipient(id interface{}) ([]FriendInvite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var invites []FriendInvite
	filter := bson.D{{Key: "recipient_id", Value: id}}
	cur, err := r.db.Collection(InviteCollection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var f FriendInvite
		invites = append(invites, f)
	}
	return invites, nil
}

func (r repo) DeleteInvite(id interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: id}}
	res, err := r.db.Collection(InviteCollection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount != 1 {
		return fmt.Errorf("no documents successfully deleted")
	}
	return nil
}

//create a new Friendship document and adds it to the collection
func (r repo) CreateFriendship(f Friendship) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := r.db.Collection(FriendshipCollection).InsertOne(ctx, f)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

//retrieves a Friendship document associated with 2 user IDs,  A and B
func (r repo) Friendship(userA, userB interface{}) (Friendship, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var f Friendship
	filter := bson.D{{Key: "users", Value: bson.D{{Key: "$all", Value: []interface{}{userA, userB}}}}}
	err := r.db.Collection(FriendshipCollection).FindOne(ctx, filter).Decode(&f)
	if err != nil {
		return Friendship{}, err
	}
	return f, nil
}

//retrieves all the Friendship documents associated with a user id
func (r repo) Friendships(userID interface{}) ([]Friendship, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var fSlice []Friendship
	filter := bson.D{{Key: "users", Value: bson.D{{Key: "$all", Value: []interface{}{userID}}}}}
	cur, err := r.db.Collection(FriendshipCollection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var f Friendship
		cur.Decode(&f)
		fSlice = append(fSlice, f)
	}
	if cur.Err() != nil {
		return nil, fmt.Errorf("failed to get friendships: %w", cur.Err())
	}
	return fSlice, nil
}

func (r repo) DeleteFriendship(id interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: id}}
	res, err := r.db.Collection(FriendshipCollection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount != 1 {
		return fmt.Errorf("no documents successfully deleted")
	}
	return nil
}

func NewRepository(db *mongo.Database) Repository {
	return repo{db: db}
}
