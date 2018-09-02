package repository

import "gopkg.in/mgo.v2"

type AbstractRepository struct {
	database *mgo.Database
	table    *mgo.Collection
}
