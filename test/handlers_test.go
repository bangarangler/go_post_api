package test

import (
	// "log"
	"testing"
)

func TestWorking(t *testing.T) {
	input := "Testing"
	expected := "Testing"

	if input != expected {
		t.Errorf("\n Word %s expected to be 'Testing', but is %s \n", input, input)
	}
}

// func ensureTableExists() {
// 	if _, err := a.DB.Exec(createSchema); err != nil {
// 		log.Fatal(err)
// 	}
// }

// const createSchema = `
// 	CREATE TABLE IF NOT EXISTS posts
// 	(
// 			id SERIAL PRIMARY KEY,
// 			title TEXT,
// 			content TEXT,
// 			author TEXT
// 	)
// `
//
// var insertPostSchema = `
// 	INSERT INTO posts(title, content, author) VALUES($1,$2,$3) RETURNING id
// `

func TestGetPostHandler(t *testing.T) {}

//
// func TestCreatePostHanndler(t *testing.T) {}
