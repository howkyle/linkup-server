package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var BadRequestBody = errors.New("unable to decode request body")
var SignupFailure = errors.New("failed to create user")

func SignupHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var us UserSignup

		err := json.NewDecoder(r.Body).Decode(&us)
		if err != nil {
			log.Println(err)
			http.Error(w, BadRequestBody.Error(), http.StatusBadRequest)
			return
		}

		_, err = s.Register(us.User())
		if err != nil {
			http.Error(w, SignupFailure.Error(), http.StatusInternalServerError)
			return
		}
	}
}
