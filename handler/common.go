package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
)

type HandlerFunc func(*sql.DB, http.ResponseWriter, *http.Request)
type OAuthHandlerFunc func(sessions.Store, http.ResponseWriter, *http.Request)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func getUserId(user interface{}) string {
	return user.(map[string]interface{})["sub"].(string)
}
