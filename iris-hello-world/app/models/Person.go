package models

import (
	"time"
)

type Person struct {
	AbstractModel
	Firstname  string      `json:"first_name"`
	Lastname   string      `json:"last_name"`
	Addresses  []*Address  `json:"addresses"`
	Documents  []*Document `json:"documents"`
	Age        int         `json:"age"`
	InsertedAt time.Time   `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time   `json:"last_update" bson:"last_update"`
}
