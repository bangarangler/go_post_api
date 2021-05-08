package database

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type PostDB interface {
	Open() error
	Close() error
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", "")
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")
	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return nil
}
