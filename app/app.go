package app

import (
	"github.com/bangarangler/go_post_api/app/database"

	"github.com/gorilla/mux"
)

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
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/posts", a.GetPostHandler()).Methods("GET")
}
