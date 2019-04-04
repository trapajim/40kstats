package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/trapajim/rest/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

type report struct {
	ListID       int    `json:"list_id"`
	UserFaction  string `json:"user_faction"`
	EnemyFaction string `json:"enemy_faction"`
	EnemyList    string `json:"enemy_list"`
	GameMode     int    `json:"game_mode"`
	Win          bool   `json:"win"`
	PlayerScore  int    `json:"player_score"`
	EnemyScore   int    `json:"enemy_score"`
}

// AddReport adds a new battle report to
func AddReport(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	userID := null.StringFrom(getUserId(r.Context().Value("user")))
	defer r.Body.Close()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	var battleReport report
	err = json.Unmarshal(b, &battleReport)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	faction := null.StringFrom("")
	if battleReport.ListID != 0 {
		list, _ := models.FindArmyList(r.Context(), db, battleReport.ListID)
		if list.UserID == userID {
			faction = list.Faction
		}
	} else {
		faction = null.StringFrom(battleReport.UserFaction)
	}
	var newBattleReport models.Battlereport
	newBattleReport.UserFaction = faction
	newBattleReport.ListID = null.IntFrom(battleReport.ListID)
	newBattleReport.Win = null.BoolFrom(battleReport.Win)
	newBattleReport.UserID = userID
	newBattleReport.EnemyFaction = null.StringFrom(battleReport.EnemyFaction)
	newBattleReport.EnemyList = null.StringFrom(battleReport.EnemyList)
	newBattleReport.EnemyScore = null.IntFrom(battleReport.EnemyScore)
	newBattleReport.PlayerScore = null.IntFrom(battleReport.PlayerScore)
	newBattleReport.GameMode = null.IntFrom(battleReport.GameMode)
	err = newBattleReport.Insert(r.Context(), db, boil.Infer())
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
}

// ListReports lists all battlereports
func ListReports(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID := null.StringFrom(getUserId(r.Context().Value("user")))
	lists, err := models.Battlereports(
		models.BattlereportWhere.UserID.EQ(userID),
		OrderBy(models.BattlereportColumns.CreatedAt)).All(r.Context(), db)

	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, lists)
}

// DeleteBattleReport removes an report from the database
func DeleteBattleReport(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	userID := null.StringFrom(getUserId(r.Context().Value("user")))
	list, err := models.FindBattlereport(r.Context(), db, id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	if list.UserID != userID {
		respondError(w, 403, "You don't have permission to change this list")
		return
	}
	list.Delete(r.Context(), db)
}

// UpdateBattleReport updates a battlereport
func UpdateBattleReport(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	userID := null.StringFrom(getUserId(r.Context().Value("user")))
	defer r.Body.Close()
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	var battleReport report
	err = json.Unmarshal(b, &battleReport)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	report, err := models.
		FindBattlereport(r.Context(), db, id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	if report.UserID != userID {
		respondError(w, 403, "You don't have permission to change this report")
		return
	}
	if battleReport.EnemyFaction != "" {
		report.EnemyFaction = null.StringFrom(battleReport.EnemyFaction)
	}
	if battleReport.EnemyList != "" {
		report.EnemyList = null.StringFrom(battleReport.EnemyList)
	}
	if battleReport.EnemyScore != 0 {
		report.EnemyScore = null.IntFrom(battleReport.EnemyScore)
	}
	if battleReport.GameMode != 0 {
		report.GameMode = null.IntFrom(battleReport.GameMode)
	}
	if battleReport.ListID != 0 {
		report.ListID = null.IntFrom(battleReport.ListID)
	}
	if battleReport.PlayerScore != 0 {
		report.PlayerScore = null.IntFrom(battleReport.PlayerScore)
	}
	if battleReport.UserFaction != "" {
		report.UserFaction = null.StringFrom(battleReport.UserFaction)
	}
	report.Win = null.BoolFrom(battleReport.Win)

	report.Update(r.Context(), db, boil.Infer())
}

type ShowBattlereportReturn struct {
	Report models.Battlereport
	List   string
}

// ShowBattleReport shows a battlereport by id
func ShowBattleReport(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID := null.StringFrom(getUserId(r.Context().Value("user")))
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	report, err := models.
		FindBattlereport(r.Context(), db, id)
	if err != nil {
		respondError(w, 404, err.Error())
		return
	}

	if report.UserID != userID {
		respondError(w, 403, "You can not view this report")
		return
	}
	list := ""
	if report.ListID.Int > 0 {
		result, err := models.FindArmyList(r.Context(), db, report.ListID.Int)
		if err == nil {
			list = result.List.String
		}
	}
	result := ShowBattlereportReturn{*report, list}
	respondJSON(w, 200, result)
}
