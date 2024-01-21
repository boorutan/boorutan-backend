package user

import (
	"applemango/boorutan/backend/db/sqlite3"
	sql "database/sql"
	"errors"
)

type User struct {
	Id string `json:"id"`
}

func CreateUser(id string) User {
	_, _ = sqlite3.DB.Exec("INSERT INTO user ( id ) VALUES ( ? )", id)
	return User{Id: id}
}

func GetUser(id string) User {
	row := sqlite3.DB.QueryRow("SELECT id FROM user WHERE id = ?", id)
	var user User
	err := row.Scan(&user.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return CreateUser(id)
	}
	if err != nil {
		panic(err)
	}
	return user
}
