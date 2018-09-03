package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	IModel `json:"-" form:"-"`
	AbstractModel
	Login    string  `json:"login"`
	Person   *Person `json:"person"`
	Password string  `json:"password" form:"-"`
}

func (*User) GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}
