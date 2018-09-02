package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"gopkg.in/mgo.v2"
	"novaxs.com/go-training/iris-hello-world/src/utils/database"
)

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	_, session, err := database.Connect()
	//Instanciando a sessão do mongo
	defer database.Close(*session)

	index := mgo.Index{
		Key:        []string{"login"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)

	if err != nil {
		panic(err)
	}

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
