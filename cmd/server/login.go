package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/GF420/go-chat/cmd/helpers"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, req *http.Request) (*Login, error) {
	var login Login

	switch req.Method {
	case http.MethodPost:
		var mr *helpers.MalformedRequest
		err := helpers.DecodeJSONBody(w, req, &login)
		if err != nil {
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return nil, err
		}

		if login.Username == "" || login.Password == "" {
			var m helpers.MalformedRequest
			m.Msg = http.StatusText(http.StatusBadRequest)
			m.Status = http.StatusBadRequest
			json.NewEncoder(w).Encode(m)
		} else {
			b, err := helpers.HashPassword(login.Password)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			login.Password = b

			json.NewEncoder(w).Encode(login)
		}

	default:
		var mr helpers.MalformedRequest
		mr.Msg = http.StatusText(http.StatusMethodNotAllowed)
		mr.Status = http.StatusMethodNotAllowed
		json.NewEncoder(w).Encode(mr)
	}

	return &login, nil
}

func LogoutHandler(w http.ResponseWriter, req *http.Request) {

}
