package app

import (
	"fmt"
	"github.com/bangarangler/go_post_api/app/models"
	"log"
	"net/http"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Post API")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapPostToJSON(p)
		sendResponse(w, r, resp, http.StatusOK)
	}
}
