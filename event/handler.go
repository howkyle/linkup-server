package event

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/howkyle/linkup-server/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//used when creating a new event
type CreateEvent struct {
	Title        string    `json:"title" validate:"required"`
	Latitude     int       `json:"latitude" validate:"required"`
	Longitude    int       `json:"longitude" validate:"required"`
	LocationName string    `json:"location_name" validate:"required"`
	Time         time.Time `json:"time" validate:"required"`
}

//maps usecase to an event
func (c CreateEvent) Event(userID primitive.ObjectID) Event {
	return Event{UserID: userID, Title: c.Title, Time: c.Time, Location: Location{Latitude: c.Latitude, Longitude: c.Longitude, LocationName: c.LocationName}}
}

//handles the creation of a new event
func NewEventHandler(s Service, v validation.Validator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sub := r.Context().Value("sub")
		userid, err := primitive.ObjectIDFromHex(fmt.Sprint(sub))
		if err != nil {
			log.Println(err)
			http.Error(w, "id conversion failed", http.StatusInternalServerError)
			return
		}

		var c CreateEvent

		err = decodeAndValidate(r.Body, &c, v)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		eid, err := s.CreateEvent(c.Event(userid))
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("failed to create event: %v", err), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(eid)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}

	}
}

func decodeAndValidate(body io.ReadCloser, target interface{}, v validation.Validator) error {
	err := json.NewDecoder(body).Decode(target)
	if err != nil {
		log.Println(err)
		return err
	}

	err = v.ValidateStruct(target)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
