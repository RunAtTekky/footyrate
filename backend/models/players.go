package models

import (
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
)

type Players struct {
	Player_List []Player
	DB          *mongo.Database
}
