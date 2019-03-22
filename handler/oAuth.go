package handler

import (
	"context"
	"crypto/rand"
	_ "crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/sessions"
	"github.com/trapajim/rest/api/config"
	"golang.org/x/oauth2"
)

// OAuthCallback handles the oAuth response
func OAuthCallback(store sessions.Store, w http.ResponseWriter, r *http.Request) {

	domain := "40kstats.auth0.com"

	conf := getOAuthConf()

	state := r.URL.Query().Get("state")
	session, err := store.Get(r, "state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if state != session.Values["state"] {
		http.Error(w, "Invalid state parameter", http.StatusInternalServerError)
		return
	}

	code := r.URL.Query().Get("code")

	token, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the userInfo
	client := conf.Client(context.TODO(), token)
	resp, err := client.Get("https://" + domain + "/userinfo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var profile map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err = store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = token.Extra("id_token")
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to logged in page
	http.Redirect(w, r, "/user", http.StatusSeeOther)

}

//Login handles the oAuth login
func Login(store sessions.Store, w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	domain := "40kstats.auth0.com"
	aud := ""
	conf := getOAuthConf()

	if aud == "" {
		aud = "https://" + domain + "/userinfo"
	}

	// Generate random state
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)

	session, err := store.Get(r, "state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	audience := oauth2.SetAuthURLParam("audience", aud)
	url := conf.AuthCodeURL(state, audience)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Logout handles oAuth logout
func Logout(store sessions.Store, w http.ResponseWriter, r *http.Request) {
	oAuthConf := config.GetConfig().OAuth
	domain := oAuthConf.Domain

	var URL *url.URL
	URL, err := url.Parse("https://" + domain)

	if err != nil {
		panic("boom")
	}
	URL.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", "http://localhost:3000")
	parameters.Add("client_id", oAuthConf.ClientId)
	URL.RawQuery = parameters.Encode()

	http.Redirect(w, r, URL.String(), http.StatusTemporaryRedirect)
}

func getOAuthConf() *oauth2.Config {
	oAuthConf := config.GetConfig().OAuth
	domain := oAuthConf.Domain

	return &oauth2.Config{
		ClientID:     oAuthConf.Id,
		ClientSecret: oAuthConf.Secret,
		RedirectURL:  oAuthConf.Callback,
		Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

}

func User(store sessions.Store, w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(session)
	renderTemplate(w, "user", session.Values["profile"])
}
