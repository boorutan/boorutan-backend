package sqlite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB = NewClient()

func Execute(query string) (any, error) {
	res, err := DB.Exec(query)
	return res, err
}

func CreateSchema() {
	Execute(`DROP TABLE IF EXISTS user`)
	Execute(
		`CREATE TABLE IF NOT EXISTS user (
			
		)`,
	)

	Execute(`DROP TABLE IF EXISTS booru`)
	Execute(
		`CREATE TABLE IF NOT EXISTS booru (
			booru_url TEXT NOT NULL,
			booru_type INTEGER NOT NULL
		)`,
	)
}

func NewClient() *sql.DB {
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		panic("Could not open database")
	}
	return db
}
