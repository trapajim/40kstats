package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/trapajim/rest/api"
	"github.com/trapajim/rest/api/config"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	boil.DebugMode = true
	config := config.GetConfig()
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	_, err3 := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err3 != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	app := &api.App{}
	app.Init(config)
	app.Run(":" + port)
}
