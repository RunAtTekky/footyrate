package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Result struct {
	Winner_ID primitive.ObjectID `bson:"winner_ID" json:"winner_ID"`
	Loser_ID  primitive.ObjectID `bson:"loser_ID" json:"loser_ID"`
}
