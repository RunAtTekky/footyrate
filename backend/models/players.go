package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Players struct {
	Images []Image
	DB     *sql.DB
}
