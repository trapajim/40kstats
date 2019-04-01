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
	bt := models.Battlereport{UserFaction: null.StringFrom("Necrons"), UserID: null.StringFrom(uid), Win: null.BoolFrom(true)}
	bt.Insert(context.Background(), a.DB, boil.Infer())
	return bt
}

func checkWrongUserBattleReport(method, path string, payload []byte, t *testing.T) {
	bt := models.Battlereport{UserFaction: null.StringFrom("Necrons"), UserID: null.StringFrom("notme")}
	bt.Insert(context.Background(), a.DB, boil.Infer())
	path = fmt.Sprintf(path, bt.ID)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(payload))
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusForbidden, response.Code, "The status code should be Forbidden")

}

func checkInvalidParameter(method, path string, payload []byte, t *testing.T) {
	path = fmt.Sprintf(path, "bbb")
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(payload))
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(t, http.StatusInternalServerError, response.Code, "The status code should be Internal Server error")
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

	checkWrongUserBattleReport("DELETE", "/v1/battlereport/%d", nil, t)
	checkInvalidParameter("DELETE", "/v1/battlereport/%s", nil, t)
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

func TestUpdateBattlereport(t *testing.T) {
	emptyTableBattleReport()
	ctx := context.Background()
	report := addBattleReport()
	assert := assert.New(t)

	path := fmt.Sprintf("/v1/battlereport/%d", report.ID)
	AuthRequired("PUT", path, t)

	payload := []byte(`{"win": false}`)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(payload))
	req.Header.Add("authorization", fmt.Sprintf("%s %s", tok.TokenType, tok.AccessToken))
	response := executeRequest(req)
	assert.Equal(http.StatusOK, response.Code, "The status code should be OK")

	updated, _ := models.Battlereports().One(ctx, a.DB)
	assert.Equal(null.BoolFrom(false), updated.Win, "Expected win to be false")
	assert.Equal(null.StringFrom("Necrons"), updated.UserFaction, "Expected userfaction to be unchanged")

	checkWrongUserBattleReport("PUT", "/v1/battlereport/%d", payload, t)
	checkInvalidParameter("PUT", "/v1/battlereport/%s", payload, t)
}
