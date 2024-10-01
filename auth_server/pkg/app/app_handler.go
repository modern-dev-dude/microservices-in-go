package app

import (
	"auth_server/pkg/dto"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		log.Println("Error while decoding login: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := h.service.Login(loginReq)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintf(w, *token)
		}
	}
}
