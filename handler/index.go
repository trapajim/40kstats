package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/volatiletech/null"

	"github.com/trapajim/rest/service"
)

type dashboard struct {
	PlayedMonth   int64                 `json:"played_month"`
	PlayedFaction chart                 `json:"played_faction"`
	WinStats      service.WinStatistics `json:"win_stats"`
}
type chart struct {
	Labels []string `json:"labels"`
	Series []int    `json:"series"`
}
type compare struct {
	Played       []int64  `json:"games_played"`
	WinStats     []int64  `json:"win_stats"`
	BestMatchup  []string `json:"best_matchup"`
	WorstMatchup []string `json:"worst_matchup"`
}

// Index inits the dashboard data
func Index(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// get user

	c, _ := service.GamesPlayedPerMonth(r.Context(), getUserId(r.Context().Value("user")), db)
	f, _ := service.GamesPlayedFaction(r.Context(), getUserId(r.Context().Value("user")), db)
	wP, _ := service.WinPercentage(r.Context(), getUserId(r.Context().Value("user")), db)
	labels := make([]string, len(f))
	series := make([]int, len(f))

	for i, element := range f {
		labels[i] = element.Faction
		series[i] = element.FactionSum
	}
	playedFactionChart := chart{Labels: labels, Series: series}
	dsb := dashboard{PlayedMonth: c, PlayedFaction: playedFactionChart, WinStats: wP}

	//fmt.Println(c)
	respondJSON(w, 200, dsb)
}

// Comparison compares two lists with each
func Comparison(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idsList := vars["ids"]
	ids := strings.Split(idsList, ",")
	played := make([]int64, len(ids))
	percentage := make([]int64, len(ids))
	bestMatchup := make([]string, len(ids))
	worstMatchup := make([]string, len(ids))
	userID := getUserId(r.Context().Value("user"))
	for i, id := range ids {
		intID, _ := strconv.Atoi(id)
		res, _ := service.GamesPlayedList(r.Context(),
			userID,
			null.IntFrom(intID), db)

		res2, _ := service.WinPercentageList(r.Context(),
			userID,
			null.IntFrom(intID), db)
		bestM, _ := service.GetMatchupStat(r.Context(), userID, null.IntFrom(intID), null.BoolFrom(true), db)
		worstM, _ := service.GetMatchupStat(r.Context(), userID, null.IntFrom(intID), null.BoolFrom(false), db)
		played[i] = res
		percentage[i] = res2.Percentage
		bestMatchup[i] = bestM.Faction
		worstMatchup[i] = worstM.Faction

	}

	ret := compare{Played: played, WinStats: percentage, BestMatchup: bestMatchup, WorstMatchup: worstMatchup}
	respondJSON(w, 200, ret)
}
