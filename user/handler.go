package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/howkyle/linkup-server/validation"
)

var BadRequestBody = errors.New("unable to decode request body")
var SignupFailure = errors.New("failed to create user")
var LoginFailure = errors.New("failed to authenticate")

type UserSignup struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u UserSignup) User() User {
	return User{Username: u.Username, Email: u.Email, Password: u.Password}
}

func SignupHandler(s Service, v validation.Validator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var us UserSignup

		err := decodeAndValidate(r.Body, &us, v)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Errorf("%v: %w", BadRequestBody.Error(), err).Error(), http.StatusBadRequest)
			return
		}

		_, err = s.Register(us.User())
		if err != nil {
			http.Error(w, fmt.Errorf("%v: %w", SignupFailure, err).Error(), http.StatusInternalServerError)
			return
		}
	}
}

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (ul UserLogin) User() User {
	return User{Username: ul.Username, Password: ul.Password}
}

func LoginHandler(s Service, v validation.Validator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ul UserLogin

		err := decodeAndValidate(r.Body, &ul, v)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Errorf("%v: %w", BadRequestBody.Error(), err).Error(), http.StatusBadRequest)
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
