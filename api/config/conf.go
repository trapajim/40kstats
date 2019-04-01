package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DB    *DBConfig
	OAuth *OAuth
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}
type OAuth struct {
	Id       string
	Domain   string
	Callback string
	Secret   string
	Audience string
	ClientId string
}

type Faction struct {
	ID   int
	Name string
}

type Detachments struct {
	Name string
	CP   int
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	env := os.Getenv("ENV")

	if env == "" {
		log.Fatal("$ENV must be set")
	}

	return &Config{
		DB: &DBConfig{
			Host:     viper.GetString(env + ".host"),
			Port:     viper.GetString(env + ".port"),
			Dialect:  "postgres",
			Username: viper.GetString(env + ".user"),
			Password: viper.GetString(env + ".pass"),
			Name:     viper.GetString(env + ".dbname"),
			Charset:  "utf8",
		},
		OAuth: &OAuth{
			Id:       viper.GetString(env + ".oauthId"),
			Domain:   viper.GetString(env + ".oauthDomain"),
			Callback: viper.GetString(env + ".oauthCallback"),
			Secret:   viper.GetString(env + ".oauthSecret"),
			Audience: viper.GetString(env + ".oauthAudience"),
			ClientId: viper.GetString(env + ".oauthClientId"),
		},
	}
}

func GetFactions() map[string][]Faction {
	elements := map[string][]Faction{}
	aeldari := []Faction{{ID: 01, Name: "Craftworlds"},
		{ID: 02, Name: "Drukhari"},
		{ID: 03, Name: "Harlequins"}, {ID: 04, Name: "Ynnari"}}
	chaos := []Faction{{ID: 11, Name: "Chaos Space Marines"},
		{ID: 12, Name: "Daemomns"},
		{ID: 13, Name: "Dark Mechanicus"},
		{ID: 14, Name: "Death Guard"},
		{ID: 15, Name: "Renegade Knights"},
		{ID: 16, Name: "Thousand Sons"}}
	imperium := []Faction{{ID: 21, Name: "Adepta Soritas"},
		{ID: 22, Name: "Adeptus Custodes"},
		{ID: 23, Name: "Adeptus Mechanicus"},
		{ID: 24, Name: "Adeptus Titanicus"},
		{ID: 25, Name: "Astra Militarum"},
		{ID: 26, Name: "Blood Angels"},
		{ID: 27, Name: "Dark Angels"},
		{ID: 28, Name: "Deathwatch"},
		{ID: 29, Name: "Grey Knights"},
		{ID: 210, Name: "Sisters of Silence"},
		{ID: 211, Name: "Space Marines"},
		{ID: 212, Name: "Space Wolves"},
	}

	elements["Aeldari"] = aeldari
	elements["Chaos"] = chaos
	elements["Imperium"] = imperium
	elements["Orks"] = []Faction{{ID: 31, Name: "Orks"}}
	elements["T'au Empire"] = []Faction{{ID: 41, Name: "T'au Empire"}}
	elements["Tyranids"] = []Faction{{ID: 51, Name: "Tyranids"}}
	elements["Gsc"] = []Faction{{ID: 61, Name: "Genestealer Cults"}}
	return elements
}

// GetDetachments returns the detachments with its cp costs
func GetDetachments() []Detachments {
	// Air Wing 1cp
	detachments := []Detachments{{"Air Wing", 1}, {"Auxiliary Support", -1},
		{"Battalion", 5}, {"Brigade", 12}, {"Fortification", 0}, {"Outrider", 1},
		{"Patrol", 0}, {"Planetstrike Attacker", 5}, {"Planetstrike Defender", 5},
		{"Spearhead", 1}, {"Stronghold Assault Attacker", 5}, {"Stronghold Assault Defender", 5},
		{"Super-Heavy Auxiliary", 0}, {"Super-Heavy", 3}, {"Supreme Command", 1}, {"Vanguard", 1}}

	return detachments
}
