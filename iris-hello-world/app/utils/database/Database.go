package database

import (
	"gopkg.in/mgo.v2"
	"novaxs.com/go-training/iris-hello-world/app/utils"
)

func Connect() (*mgo.Database, mgo.Session, error) {
	session, err := mgo.Dial(utils.DB_LINK)
	session.SetMode(mgo.Monotonic, true)

	return session.DB(utils.DB_DATABASE), *session, err
}

func Close(session mgo.Session) {
	session.Close()
}
