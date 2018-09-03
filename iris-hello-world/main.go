package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"novaxs.com/go-training/iris-hello-world/app/utils/database"
	"novaxs.com/go-training/iris-hello-world/routes"
)

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	//app.Use(recover.New())
	app.Use(logger.New())
	_, session, err := database.Connect()
	//Chama as rotas da aplicação
	routes.Index(app)

	if err != nil {
		panic(err)
	}
	session.Close()
	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
