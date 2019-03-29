package service

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/trapajim/rest/api/config"
	"github.com/trapajim/rest/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

var uid string
var db *sql.DB

func TestMain(m *testing.M) {
	config := config.GetConfig()
	uid = "uid1"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.Name)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	ensureTableExists(config)
	clearDb()
	setupDb()
	code := m.Run()
	os.Exit(code)

}
func ensureTableExists(config *config.Config) {
	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations",
	}
	_, err3 := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err3 != nil {
		fmt.Println(err3)
	}
}
func setupDb() {
	ctx := context.Background()
	var list models.ArmyList
	list.Faction = null.StringFrom("Necrons")
	list.UserID = null.StringFrom(uid)
	list.Insert(ctx, db, boil.Infer())
	var list2 models.ArmyList
	list2.Faction = null.StringFrom("Orks")
	list2.UserID = null.StringFrom(uid)
	list2.Insert(ctx, db, boil.Infer())
	bt := models.Battlereport{UserFaction: list.Faction, ListID: null.IntFrom(list.ID), EnemyFaction: null.StringFrom("Adeptus Custodes"), Win: null.BoolFrom(true), UserID: null.StringFrom(uid)}
	bt.Insert(context.Background(), db, boil.Infer())
	bt2 := models.Battlereport{UserFaction: list.Faction, ListID: null.IntFrom(list.ID),
		EnemyFaction: null.StringFrom("Adeptus Custodes"),
		Win:          null.BoolFrom(true), UserID: null.StringFrom(uid)}
	bt2.Insert(context.Background(), db, boil.Infer())
	bt3 := models.Battlereport{UserFaction: list.Faction, ListID: null.IntFrom(list.ID),
		EnemyFaction: null.StringFrom("Adeptus Mechanicus"),
		Win:          null.BoolFrom(false), UserID: null.StringFrom(uid)}
	bt3.Insert(context.Background(), db, boil.Infer())
	bt4 := models.Battlereport{UserFaction: list2.Faction, ListID: null.IntFrom(list2.ID),
		EnemyFaction: null.StringFrom("Adeptus Mechanicus"),
		Win:          null.BoolFrom(false), UserID: null.StringFrom(uid)}
	bt4.Insert(context.Background(), db, boil.Infer())

}

func clearDb() {
	models.Battlereports().DeleteAll(context.Background(), db)
	models.ArmyLists().DeleteAll(context.Background(), db)
}

func TestGamesPlayedPerMonth(t *testing.T) {

	type args struct {
		ctx    context.Context
		userID string
		db     *sql.DB
	}
	arg := args{context.Background(), uid, db}

	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"should return the amount of games played of the current month", arg, int64(4), false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GamesPlayedPerMonth(tt.args.ctx, tt.args.userID, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GamesPlayedPerMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GamesPlayedPerMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGamesPlayedFaction(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
		db     *sql.DB
	}
	arg := args{context.Background(), uid, db}
	tests := []struct {
		name    string
		args    args
		want    []GamesPlayedFactionStruct
		wantErr bool
	}{
		{"tests games played per faction", arg, []GamesPlayedFactionStruct{{FactionSum: 3, Faction: "Necrons"}, {FactionSum: 1, Faction: "Orks"}}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GamesPlayedFaction(tt.args.ctx, tt.args.userID, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GamesPlayedFaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GamesPlayedFaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWinPercentage(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
		db     *sql.DB
	}
	arg := args{context.Background(), uid, db}
	tests := []struct {
		name    string
		args    args
		want    WinStatistics
		wantErr bool
	}{
		{"overall win statistics", arg, WinStatistics{Losses: 2, Wins: 2, Percentage: 50}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WinPercentage(tt.args.ctx, tt.args.userID, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("WinPercentage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WinPercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGamesPlayedList(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
		list   null.Int
		db     *sql.DB
	}
	list, _ := models.ArmyLists(models.ArmyListWhere.Faction.EQ(null.StringFrom("Necrons"))).One(context.Background(), db)
	arg := args{context.Background(), uid, null.IntFrom(list.ID), db}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"should return the number of games played for a list", arg, int64(3), false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GamesPlayedList(tt.args.ctx, tt.args.userID, tt.args.list, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GamesPlayedList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GamesPlayedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWinPercentageList(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
		list   null.Int
		db     *sql.DB
	}
	list, _ := models.ArmyLists(models.ArmyListWhere.Faction.EQ(null.StringFrom("Necrons"))).One(context.Background(), db)
	arg := args{context.Background(), uid, null.IntFrom(list.ID), db}
	tests := []struct {
		name    string
		args    args
		want    WinStatistics
		wantErr bool
	}{
		{"should return the win stats for a given list", arg, WinStatistics{Losses: 1, Wins: 2, Percentage: 66}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WinPercentageList(tt.args.ctx, tt.args.userID, tt.args.list, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("WinPercentageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WinPercentageList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMatchupStat(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
		list   null.Int
		win    null.Bool
		db     *sql.DB
	}
	list, _ := models.ArmyLists(models.ArmyListWhere.Faction.EQ(null.StringFrom("Necrons"))).One(context.Background(), db)
	arg := args{context.Background(), uid, null.IntFrom(list.ID), null.BoolFrom(false), db}
	arg2 := args{context.Background(), uid, null.IntFrom(list.ID), null.BoolFrom(true), db}
	tests := []struct {
		name    string
		args    args
		want    Matchup
		wantErr bool
	}{
		{"should give the worst matchups", arg, Matchup{Faction: "Adeptus Mechanicus", Sum: 1}, false},
		{"should give the best matchups", arg2, Matchup{Faction: "Adeptus Custodes", Sum: 2}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMatchupStat(tt.args.ctx, tt.args.userID, tt.args.list, tt.args.win, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchupStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatchupStat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcPercentage(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"should return the percent of two ints", args{5, 5}, 50},
		{"should fallback if two 0 are parameters", args{0, 0}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcPercentage(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("calcPercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}
