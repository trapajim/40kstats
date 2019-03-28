package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trapajim/rest/api/config"

	"github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
)

// User is a struct which keeps the username and id
type User struct {
	Nickname string
	Sub      string
}

// AuthRequired is a middleware to check if the user is authenticated
func AuthRequired(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.GetConfig().OAuth
		var auth0Domain = "https://" + conf.Domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{conf.Audience}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		out := make(map[string]interface{})
		token.UnsafeClaimsWithoutVerification(&out)
		inner.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", out)))
	})
}

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
