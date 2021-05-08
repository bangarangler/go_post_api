package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")

	pg.MustExec(createSchema)

	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
