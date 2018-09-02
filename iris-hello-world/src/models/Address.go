package models

import (
	"time"
)

//Type => se é residencial, comercial etc.
type Address struct {
	//ID         bson.ObjectId `bson:"_id,omitempty"`
	Street     string    `json:"street"`
	Number     string    `json:"number"`
	Type       string    `json:"type"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	InsertedAt time.Time `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time `json:"last_update" bson:"last_update"`
}
