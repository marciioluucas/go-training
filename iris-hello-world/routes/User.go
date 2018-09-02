package routes

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"gopkg.in/mgo.v2/bson"
	"novaxs.com/go-training/iris-hello-world/src/models"
	"time"
)

func UserRoutes(app *iris.Application) {
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.JSON(context.Map{"message": "Bem vindo ao micro service de usuário"})
	})

	app.Handle("GET", "/users", func(ctx context.Context) {
		results := []models.User{}
		err := c.Find(nil).All(&results)
		if err != nil {
			ctx.JSON(context.Map{"message": "Algo de errado não deu certo"})
		} else {
			fmt.Println("Results All: ", results)
		}
		ctx.JSON(context.Map{"response": results})

	})

	app.Handle("GET", "/user/{login: string}", func(ctx context.Context) {
		login := ctx.Params().Get("login")
		fmt.Println(login)
	})

	app.Handle("POST", "/user", func(ctx context.Context) {
		params := &models.User{}
		err := ctx.ReadJSON(params)
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
		} else {
			params.LastUpdate = time.Now()
			byteArray, error := models.GeneratePassword(params.Password)
			//n := bytes.IndexByte(byteArray, 0)
			params.Password = string(byteArray)
			err := c.Insert(params)
			if error != nil {
				err = error
			}
			if err != nil {
				ctx.JSON(context.Map{"response": err.Error()})
			} else {
				fmt.Println("Usuario inserido com sucesso!")
				result := models.User{}
				err = c.Find(bson.M{"login": params.Login}).One(&result)
				if err != nil {
					ctx.JSON(context.Map{"response": err.Error()})
				}
				ctx.JSON(context.Map{"response": "Usuário criado com sucesso!", "message": result})
			}
		}
	})
}
