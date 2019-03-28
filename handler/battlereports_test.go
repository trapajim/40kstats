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

func emptyTableBattleReport() {
	models.Battlereports().DeleteAll(context.Background(), a.DB)
}

func addBattleReport() models.Battlereport {
	bt := models.Battlereport{UserFaction: null.StringFrom("Necrons"), UserID: null.StringFrom(uid)}
	bt.Insert(context.Background(), a.DB, boil.Infer())
	return bt
}

func TestListReports(t *testing.T) {
	AuthRequired("GET", "/v1/battlereport", t)
	req, _ := http.NewRequest("GET", "/v1/battlereport", nil)
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code, "The status code should be OK")
}

func TestDeleteBattleReport(t *testing.T) {
	emptyTableBattleReport()
	rp := addBattleReport()
	assert := assert.New(t)
	ctx := context.Background()
	len, _ := models.Battlereports().Count(ctx, a.DB)
	assert.Equal(int64(1), len, "Battlereports should have 1 entry")
	AuthRequired("DELETE", "/v1/battlereport/1", t)
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/battlereport/%d", rp.ID), nil)
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(http.StatusOK, response.Code, "The status code should be OK")
	len, _ = models.Battlereports().Count(ctx, a.DB)
	assert.Equal(int64(0), len, "Battlereports should be empty")
}

func TestAddReport(t *testing.T) {
	emptyTableBattleReport()
	ctx := context.Background()
	var list models.ArmyList
	list.Faction = null.StringFrom("Necrons")
	list.UserID = null.StringFrom(uid)
	list.Insert(ctx, a.DB, boil.Infer())

	assert := assert.New(t)
	AuthRequired("POST", "/v1/battlereport", t)
	len, _ := models.Battlereports().Count(ctx, a.DB)
	assert.Equal(int64(0), len, "Battlereports should have 0 entry")
	payload := []byte(fmt.Sprintf(`{"list_id": %d,"enemy_faction":"necrons","user_faction":"necrons", "win": true}`, list.ID))
	req, _ := http.NewRequest("POST", "/v1/battlereport", bytes.NewBuffer(payload))
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(http.StatusOK, response.Code, "The status code should be OK")
	len, _ = models.Battlereports().Count(ctx, a.DB)
	assert.Equal(int64(1), len, "Battlereports should be empty")
}