package event

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//handles the creation of a new event
func NewEventHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.Context().Value("sub")
		var c CreateEvent
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}

		c.UserID, err = primitive.ObjectIDFromHex(fmt.Sprint(userid))
		if err != nil {
			log.Println(err)
			http.Error(w, "id conversion failed", http.StatusInternalServerError)
			return
		}

		eid, err := s.CreateEvent(c)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to create event", http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(eid)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}

	}
}
