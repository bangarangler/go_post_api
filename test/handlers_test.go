package test

import (
	"log"
	"testing"

	"github.com/bangarangler/go_post_api/app"
	"github.com/bangarangler/go_post_api/database"
)

func ensureTableExists() {
	if _, err := a.DB.Exec(createSchema); err != nil {
		log.Fatal(err)
	}
}

func TestGetPostHandler(t *testing.T) {}

func TestCreatePostHanndler(t *testing.T) {}
