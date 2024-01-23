package sqlite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB = NewClient()
var TagDB *sql.DB = (func() *sql.DB {
	db, err := sql.Open("sqlite3", "/Users/apple/Documents/docs/www/full/boorutan/booru-japanese-tag/app.db")
	if err != nil {
		panic("Could not open database")
	}
	return db
})()

func Execute(query string) (any, error) {
	res, err := DB.Exec(query)
	return res, err
}

func CreateSchema(query string) {
	_, _ = Execute(query)
}

func InitDB() {
	CreateSchema(`
		DROP TABLE IF EXISTS user
	`)
	CreateSchema(`
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`)
	CreateSchema(`
		CREATE INDEX user_name_index on user ( username )	
	`)
	CreateSchema(`
		INSERT INTO user (username, password) VALUES  ( "apple", "dev" )
	`)
	CreateSchema(`
		DROP TABLE IF EXISTS like
	`)
	CreateSchema(`
		CREATE TABLE IF NOT EXISTS like (
			id INTEGER PRIMARY KEY  AUTOINCREMENT,
			booru TEXT NOT NULL,
			post_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY ( user_id ) REFERENCES user ( id ) ON DELETE CASCADE
		)
	`)
	CreateSchema(`
		CREATE INDEX user_liked_post_index on like ( user_id )
	`)
}

func NewClient() *sql.DB {
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		panic("Could not open database")
	}
	return db
}
