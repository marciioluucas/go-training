package models

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Login      string        `json:"login"`
	Person     *Person       `json:"person"`
	Password   string        `json:"-" form:"-"`
	InsertedAt time.Time     `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time     `json:"last_update" bson:"last_update"`
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}
