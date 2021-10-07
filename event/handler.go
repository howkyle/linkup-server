package event

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func NewEventHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := fmt.Sprintf("%v", r.Context().Value("sub"))

		uid, err := strconv.ParseUint(userid, 10, 64)
		if err != nil {
			http.Error(w, "failed to get user", http.StatusInternalServerError)
			return
		}

		var c CreateEvent
		err = json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}

		c.UserID = uint(uid)
		_, err = s.CreateEvent(c)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to create event", http.StatusInternalServerError)
			return
		}

	}
}
