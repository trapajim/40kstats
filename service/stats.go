package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/trapajim/rest/models"
	"github.com/volatiletech/null"
)

// Custom struct for selecting a subset of data
type GamesPlayedFactionStruct struct {
	FactionSum int    `boil:"faction_sum"`
	Faction    string `boil:"user_faction"`
}

type Matchup struct {
	Sum     int    `boil:"sum"`
	Faction string `boil:"faction"`
}

// Custom struct for selecting a subset of data
type GamesPlayed struct {
	GamesPlayed int `boil:"games_played"`
}

// WinStatistics
type WinStatistics struct {
	Percentage int64 `json:"percentage"`
	Wins       int64 `json:"wins"`
	Losses     int64 `json:"losses"`
}

// GamesPlayedPerMonth determines the games played per month for a user
func GamesPlayedPerMonth(ctx context.Context, userID string, db *sql.DB) (int64, error) {
	t := time.Now()
	firstday := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	lastday := firstday.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	count, err := models.Battlereports(
		models.BattlereportWhere.UserID.EQ(null.StringFrom(userID)),
		models.BattlereportWhere.CreatedAt.GT(firstday),
		models.BattlereportWhere.CreatedAt.LT(lastday)).Count(ctx, db)
	if err != nil {

		return 0, err
	}
	return count, nil

}

// GamesPlayedFaction determines the games played per faction for a user
func GamesPlayedFaction(ctx context.Context, userID string, db *sql.DB) ([]GamesPlayedFactionStruct, error) {
	var result []GamesPlayedFactionStruct
	err := models.Battlereports(
		qm.Select("user_faction", "count(\"user_faction\") as faction_sum"),
		models.BattlereportWhere.UserID.EQ(null.StringFrom(userID)),
		qm.GroupBy("user_faction")).Bind(ctx, db, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//WinPercentage calculates the overall win percentage
func WinPercentage(ctx context.Context, userID string, db *sql.DB) (WinStatistics, error) {
	wins, err := models.Battlereports(models.BattlereportWhere.Win.EQ(null.BoolFrom(true)),
		models.BattlereportWhere.UserID.EQ(null.StringFrom(userID))).Count(ctx, db)
	stat := WinStatistics{}
	if err != nil {
		return stat, err
	}
	loss, err := models.Battlereports(models.BattlereportWhere.Win.EQ(null.BoolFrom(false)),
		models.BattlereportWhere.UserID.EQ(null.StringFrom(userID))).Count(ctx, db)
	if err != nil {
		return stat, err
	}
	stat = WinStatistics{Wins: wins, Losses: loss, Percentage: calcPercentage(wins, loss)}
	return stat, nil
}

//GamesPlayedList the games played for one faction
func GamesPlayedList(ctx context.Context, userID string, list null.Int, db *sql.DB) (int64, error) {
	i, err := models.Battlereports(
		models.BattlereportWhere.UserID.EQ(null.StringFrom(userID)),
		models.BattlereportWhere.ListID.EQ(list),
	).Count(ctx, db)

	if err != nil {
		return 0, err
	}

	return i, nil
}

//WinPercentageList calculates the overall win percentage for a list
func WinPercentageList(ctx context.Context, userID string, list null.Int, db *sql.DB) (WinStatistics, error) {
	wins, err := models.Battlereports(models.BattlereportWhere.Win.EQ(null.BoolFrom(true)),
		models.BattlereportWhere.UserID.EQ(null.StringFrom(userID)), models.BattlereportWhere.ListID.EQ(list)).Count(ctx, db)
	stat := WinStatistics{}
	if err != nil {
		return stat, err
	}
	loss, err := models.Battlereports(models.BattlereportWhere.Win.EQ(null.BoolFrom(false)),
		models.BattlereportWhere.UserID.EQ(null.StringFrom(userID)), models.BattlereportWhere.ListID.EQ(list)).Count(ctx, db)
	if err != nil {
		return stat, err
	}

	stat = WinStatistics{Wins: wins, Losses: loss, Percentage: calcPercentage(wins, loss)}
	return stat, nil
}

func GetMatchupStat(ctx context.Context, userID string, list null.Int, win null.Bool, db *sql.DB) (Matchup, error) {
	var result Matchup
	err := models.Battlereports(
		qm.Select("enemy_faction as faction", "count(\"enemy_faction\") as sum"),
		models.BattlereportWhere.Win.EQ(win),
		models.BattlereportWhere.UserID.EQ(null.StringFrom(userID)),
		models.BattlereportWhere.ListID.EQ(list), qm.GroupBy("enemy_faction"),
		qm.OrderBy("sum DESC"),
		qm.Limit(1)).Bind(ctx, db, &result)

	if err != nil {
		return result, err
	}
	return result, nil
}

func calcPercentage(a, b int64) int64 {
	max := a + b
	if max == 0 {
		max = 1
	}
	return ((a * 100) / max)
}
