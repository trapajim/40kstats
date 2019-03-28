package handler_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/stretchr/testify/assert"
	"github.com/trapajim/rest/api"
	"github.com/trapajim/rest/api/config"
)

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

var a api.App
var tok token
var uid string

func TestMain(m *testing.M) {
	config := config.GetConfig()
	a = api.App{}
	a.Init(config)
	uid = fmt.Sprintf("%s@clients", config.OAuth.ClientId)
	ensureTableExists(config)
	getToken(config)
	code := m.Run()
	os.Exit(code)
}

func ensureTableExists(config *config.Config) {
	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations",
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	_, err3 := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err3 != nil {
		fmt.Println(err)
	}
}

func getToken(config *config.Config) {
	url := "https://" + config.OAuth.Domain + "/oauth/token"
	s := fmt.Sprintf("{\"client_id\":\"%s\",\"client_secret\":\"%s\",\"audience\":\"%s\",\"grant_type\":\"client_credentials\"}", config.OAuth.ClientId, config.OAuth.Secret, config.OAuth.Audience)
	payload := strings.NewReader(s)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &tok)
}

func AuthRequired(method, path string, t *testing.T) {
	req, _ := http.NewRequest(method, path, nil)
	response := executeRequest(req)
	assert.Equal(t, http.StatusForbidden, response.Code, "Http status code should be 403 for unauthorized users")

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}
