//handles events
package event

import "time"

type Event interface {
	//returns the location of the event
	Location() (Location, error)
	//returns the time that the event should occur
	Time() time.Time
	//adds a new participant to the event
	AddParticipant(p Participant) error
	//returns all the participants
	Participants() []Participant
}

type Location struct {
	Latitude  int
	Longitude int
	Name      string
}

type Participant interface {
}
