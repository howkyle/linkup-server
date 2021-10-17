package invitation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/howkyle/linkup-server/event"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewInvitationHandler(s Service, e event.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c CreateInvitation
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}
		event, err := e.Event(c.EventID)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to get associated event", http.StatusInternalServerError)
		}
		c.Summary = fmt.Sprintf("%v at %v, at %s", event.Title, event.LocationName, event.Time)
		id, err := s.AddInvitation(c.Invitation())
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to add invite", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(id)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func GetInvitationsHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sub := r.Context().Value("sub")
		uid, err := primitive.ObjectIDFromHex(fmt.Sprint(sub))
		if err != nil {
			log.Println(err)
			http.Error(w, "id conversion failed", http.StatusInternalServerError)
			return
		}
		i, err := s.InvitationsByUser(uid)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to retrieve invitations", http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		err = json.NewEncoder(w).Encode(i)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}

func AcceptHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sub := r.Context().Value("sub")
		uid, err := primitive.ObjectIDFromHex(fmt.Sprint(sub))
		if err != nil {
			log.Println(err)
			http.Error(w, "id conversion failed", http.StatusInternalServerError)
			return
		}

		params := mux.Vars(r)
		invitationID := params["id"]
		if invitationID == "" {
			http.Error(w, "missing param: invitation id", http.StatusBadRequest)
			return
		}

		id, err := primitive.ObjectIDFromHex(fmt.Sprintf(invitationID))
		if err != nil {
			log.Println(err)
			http.Error(w, "invalid invitation id", http.StatusBadRequest)
		}

		err = s.AcceptInvitation(uid, id)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to accept invitation", http.StatusInternalServerError)
		}
	}
}

func createResponse(body interface{}) (string, error) {
	response, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("unable to create response: %v", err)
	}
	return string(response), nil
}

func writeJSON(w http.ResponseWriter, v string) {
	w.Header().Add("content-type", "application/json")
	fmt.Fprintf(w, v)
}
