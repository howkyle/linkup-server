package invitation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewInvitationHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c CreateInvitation
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}
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

		err = json.NewEncoder(w).Encode(i)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}
