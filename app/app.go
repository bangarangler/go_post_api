package app

import (
	// "fmt"
	// "os"
	"github.com/bangarangler/go_post_api/app/database"
	"github.com/gorilla/mux"
	// "github.com/hoho/godotenv"
)

// func goDotEnvVar(key string) string {
// 	err := godotenv.Load("db.env")
// 	if err != nil {
// 		log.Fatalf("error loading env")
// 	}
// 	return os.Getenv(key)
// }

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
}
