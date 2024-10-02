package app

import (
	"encoding/json"
	"github.com/modern-dev-dude/microservices-in-go/auth_server/pkg/dto"
	"github.com/modern-dev-dude/microservices-in-go/auth_server/pkg/service"
	"log"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		log.Println("Error while decoding login: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := h.service.Login(loginReq)
		if err != nil {
			writeResponse(w, http.StatusUnauthorized, "un authorized access")
		} else {
			writeResponse(w, http.StatusOK, *token)
		}
	}
}

func (h AuthHandler) NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "handler not implemented")
}

func (h AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	for key := range r.URL.Query() {
		urlParams[key] = r.URL.Query().Get(key)
	}

	if urlParams["token"] != "" {
		err := h.service.Verify(urlParams)
		if err != nil {
			writeResponse(w, http.StatusUnauthorized, unauthorizedResponse("unauthorized"))
		} else {
			writeResponse(w, http.StatusOK, authorizedResponse())
		}
	}
}

func (h AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var refreshRequest dto.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&refreshRequest); err != nil {
		log.Println("Error while decoding refresh: " + err.Error())
		writeResponse(w, http.StatusBadRequest, "bad request")
	} else {
		token, err := h.service.Refresh(refreshRequest)
		if err != nil {
			writeResponse(w, http.StatusUnauthorized, "unauthorized access")
		} else {
			writeResponse(w, http.StatusOK, *token)
		}
	}
}

func unauthorizedResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"isAuthorized": false,
		"message":      msg,
	}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
}

func writeResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		panic(err)
	}
}
