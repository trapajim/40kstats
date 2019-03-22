package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	. "github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/gorilla/mux"
	"github.com/trapajim/rest/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

type newList struct {
	Name    string `json:"name"`
	List    string `json:"list"`
	Faction string `json:"faction"`
}

type user struct {
	sub string
}

//AddNewList adds a new army list
func AddNewList(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var msg newList
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userID := getUserId(r.Context().Value("user"))

	var newArmyList models.ArmyList
	newArmyList.Faction = null.StringFrom(msg.Faction)
	fmt.Println(msg.Faction)
	newArmyList.List = null.StringFrom(msg.List)
	newArmyList.ListName = null.StringFrom(msg.Name)
	newArmyList.UserID = null.StringFrom(userID)
	newArmyList.Insert(r.Context(), db, boil.Infer())
	fmt.Println(msg)
}

// UpdateList updates a list by ID
func UpdateList(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var msg newList
	err := json.Unmarshal(b, &msg)
	userId := null.StringFrom(getUserId(r.Context().Value("user")))
	vars := mux.Vars(r)
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}

	list, err := models.
		FindArmyList(r.Context(), db, id)

	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	if list.UserID != userId {
		respondError(w, 403, "You don't have permission to change this list")
		return
	}

	list.ListName = null.StringFrom(msg.Name)
	list.Faction = null.StringFrom(msg.Faction)
	list.List = null.StringFrom(msg.List)
	list.Update(r.Context(), db, boil.Infer())
}

// DeleteList removes a list
func DeleteList(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var msg newList
	_ = json.Unmarshal(b, &msg)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, 500, err.Error())
		return
	}
	userId := null.StringFrom(getUserId(r.Context().Value("user")))
	list, err := models.
		FindArmyList(r.Context(), db, id)

	if err != nil {
		respondError(w, 404, err.Error())
		return
	}
	if list.UserID != userId {
		respondError(w, 403, "You don't have permission to change this list")
		return
	}

	list.Delete(r.Context(), db)
}

// ShowLists lists all the army lists a user has.
func ShowLists(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID := null.StringFrom(getUserId(r.Context().Value("user")))
	lists, err := models.ArmyLists(
		models.ArmyListWhere.UserID.EQ(userID),
		OrderBy(models.ArmyListColumns.ListName)).All(r.Context(), db)

	if err != nil {
		respondError(w, 400, err.Error())
		return
	}
	respondJSON(w, 200, lists)
}
