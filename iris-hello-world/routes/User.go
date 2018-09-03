package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"novaxs.com/go-training/iris-hello-world/app/models"
	"novaxs.com/go-training/iris-hello-world/app/repository"
)

func UserRoutes(app *iris.Application) {
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.JSON(context.Map{"message": "Bem vindo ao micro service de usuário"})
	})

	app.Handle("GET", "/users", func(ctx context.Context) {
		repo := (&repository.UserRepository{}).New()
		ctx.JSON(context.Map{"users": repo.Retrieve()})
	})

	app.Handle("GET", "/user/{login: string}", func(ctx context.Context) {
		login := ctx.Params().Get("login")
		repo := (&repository.UserRepository{}).New()
		ctx.JSON(context.Map{"user": repo.Find(login)})
	})

	app.Handle("POST", "/user", func(ctx context.Context) {
		params := &models.User{}
		ctx.ReadJSON(&params)
		repo := (&repository.UserRepository{}).New().Create(*params)
		ctx.JSON(context.Map{"user": repo})
	})
}
