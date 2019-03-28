package handler_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trapajim/rest/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func TestAddNewList(t *testing.T) {
	emptyTableArmyList()

	ctx := context.Background()
	assert := assert.New(t)
	AuthRequired("POST", "/v1/list", t)
	len, _ := models.ArmyLists().Count(ctx, a.DB)

	assert.Equal(int64(0), len, "Army lists should have 0 entry")
	payload := []byte(`{"faction": "Necrons"}`)
	req, _ := http.NewRequest("POST", "/v1/list", bytes.NewBuffer(payload))
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(http.StatusOK, response.Code, "The status code should be OK")

	len, _ = models.ArmyLists().Count(ctx, a.DB)

	assert.Equal(int64(1), len, "ArmyLists should have 1 entry")

}

func emptyTableArmyList() {
	models.ArmyLists().DeleteAll(context.Background(), a.DB)
}

func TestUpdateList(t *testing.T) {
	assert := assert.New(t)
	emptyTableArmyList()
	ctx := context.Background()
	list := addArmyList()
	path := fmt.Sprintf("/v1/list/%d", list.ID)
	AuthRequired("PUT", path, t)

	payload := []byte(`{"faction": "Orks"}`)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(payload))
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(http.StatusOK, response.Code, "The status code should be OK")

	updated, _ := models.ArmyLists().One(ctx, a.DB)
	assert.Equal("Orks", updated.Faction.String, "Expected the faction to be changed")

	checkWrongUser("PUT", "/v1/list/%d", payload, t)
	checkNotFound("PUT", "/v1/list/%d", payload, t)
}

func TestDeleteList(t *testing.T) {
	emptyTableArmyList()
	assert := assert.New(t)
	ctx := context.Background()
	list := addArmyList()
	len, _ := models.ArmyLists().Count(ctx, a.DB)
	assert.Equal(int64(1), len, "Lists should have 1 entry")
	path := fmt.Sprintf("/v1/list/%d", list.ID)
	AuthRequired("DELETE", path, t)

	req, _ := http.NewRequest("DELETE", path, nil)
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(http.StatusOK, response.Code, "The status code should be OK")

	len, _ = models.ArmyLists().Count(ctx, a.DB)
	assert.Equal(int64(0), len, "ArmyLists should be empty")

	req, _ = http.NewRequest("DELETE", "/v1/list/0", nil)
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response = executeRequest(req)
	assert.Equal(http.StatusNotFound, response.Code, "The status code should be OK")

	checkWrongUser("DELETE", "/v1/list/%d", nil, t)
	checkNotFound("DELETE", "/v1/list/%d", nil, t)
}

func TestShowLists(t *testing.T) {
	AuthRequired("GET", "/v1/show-list", t)
	req, _ := http.NewRequest("GET", "/v1/show-list", nil)
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code, "The status code should be OK")
}

func checkWrongUser(method, path string, payload []byte, t *testing.T) {
	rndList := addArmyListUid("notme")
	path = fmt.Sprintf(path, rndList.ID)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(payload))
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusForbidden, response.Code, "The status code should be Forbidden")

}

func checkNotFound(method, path string, payload []byte, t *testing.T) {
	path = fmt.Sprintf(path, 0)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(payload))
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusNotFound, response.Code, "The status code should be Not found")
}

func addArmyList() models.ArmyList {
	var list models.ArmyList
	list.Faction = null.StringFrom("Necrons")
	list.UserID = null.StringFrom(uid)
	list.Insert(context.Background(), a.DB, boil.Infer())
	return list
}

func addArmyListUid(rndID string) models.ArmyList {
	var list models.ArmyList
	list.Faction = null.StringFrom("Necrons")
	list.UserID = null.StringFrom(rndID)
	list.Insert(context.Background(), a.DB, boil.Infer())
	return list
}
