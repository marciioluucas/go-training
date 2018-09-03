package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type AbstractModel struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	InsertedAt time.Time     `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time     `json:"last_update" bson:"last_update"`
}
