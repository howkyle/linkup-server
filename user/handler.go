package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var BadRequestBody = errors.New("unable to decode request body")
var SignupFailure = errors.New("failed to create user")
var LoginFailure = errors.New("failed to authenticate")

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

func LoginHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ul UserLogin

		err := json.NewDecoder(r.Body).Decode(&ul)
		if err != nil {
			log.Println(err)
			http.Error(w, BadRequestBody.Error(), http.StatusBadRequest)
			return
		}

		res, err := s.Login(ul.User())
		if err != nil {
			log.Println(err)
			http.Error(w, LoginFailure.Error(), http.StatusUnauthorized)
			return
		}

		cookie, ok := res.(http.Cookie)
		if !ok {
			http.Error(w, LoginFailure.Error(), http.StatusInternalServerError)
		}
		http.SetCookie(w, &cookie)
	}
}
