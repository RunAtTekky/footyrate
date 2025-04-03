package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	URL      string             `bson:"url" json:"url"`
	ELO      int                `bson:"elo" json:"elo"`
	K_FACTOR int                `bson:"k_factor" json:"k_factor"`
	ROUNDS   int                `bson:"rounds" json:"rounds"`
}
