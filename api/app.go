package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"encoding/gob"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/trapajim/rest/api/config"
	"github.com/trapajim/rest/handler"
	"github.com/trapajim/rest/router"
)

//App represents the App
type App struct {
	Router *mux.Router
	DB     *sql.DB
	Store  *sessions.FilesystemStore
	OAuth  *config.OAuth
}

// Init initializes the App
func (a *App) Init(config *config.Config) {

	gob.Register(map[string]interface{}{})

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter().StrictSlash(true)
	a.DB = db
	a.Store = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	a.OAuth = config.OAuth
	a.Router = r

	a.Router.Use(router.Logger)
	//a.Router.Use(router.Cors)
	a.Router.Use(router.AuthRequired)

	a.setAPIRoutes()
}

func (a *App) makeApiHandler(fn handler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(a.DB, w, r)
	}
}

func (a *App) setAPIRoutes() {
	sub := a.Router.PathPrefix("/v1").Subrouter()
	fmt.Println("Available API: ")
	for _, route := range router.GetRoutes {
		fmt.Println(route.Name, route.Method, route.Pattern)
		sub.
			HandleFunc(route.Pattern, a.makeApiHandler(route.HandlerFunc)).
			Name(route.Name).
			Methods(route.Method, "OPTIONS")
	}
}

// Run starts the app server
func (a *App) Run(host string) {

	fmt.Println("Server running at localhost", host)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001", "https://stats40k.herokuapp.com"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
	})
	handler := c.Handler(a.Router)

	log.Fatal(http.ListenAndServe(host, handler))
}
