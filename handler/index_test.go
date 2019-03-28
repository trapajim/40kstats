package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	AuthRequired("GET", "/v1/dashboard", t)
	req, _ := http.NewRequest("GET", "/v1/dashboard", nil)
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code, "The status code should be OK")
}

func TestFactions(t *testing.T) {
	AuthRequired("GET", "/v1/factions", t)
	req, _ := http.NewRequest("GET", "/v1/factions", nil)
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code, "The status code should be OK")
}

func TestComparision(t *testing.T) {
	AuthRequired("GET", "/v1/compare/1,2", t)
	req, _ := http.NewRequest("GET", "/v1/compare/1,2", nil)
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code, "The status code should be OK")
}
