package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Players struct {
	Player_List []Player
	DB          *sql.DB
}
