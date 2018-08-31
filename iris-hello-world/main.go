package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/context"
)

type User struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Firstname  string        `json:"firstname"`
	Lastname   string        `json:"lastname"`
	Age        int           `json:"age"`
	Msisdn     string        `json:"msisdn"`
	InsertedAt time.Time     `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time     `json:"last_update" bson:"last_update"`
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	//Instanciando a sessão do mongo
	session, err := mgo.Dial("127.0.0.1")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("usergo").C("profiles")

		index := mgo.Index{
		Key:        []string{"msisdn"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = c.EnsureIndex(index)
	
	if err != nil {
		panic(err)
	}

	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.JSON(context.Map{"message": "Bem vindo ao micro service de usuário"})
	})

	app.Handle("GET", "/users", func(ctx context.Context) {
		results := []User{}

		err := c.Find(nil).All(&results)
		if(err != nil) {
			ctx.JSON(context.Map{"message": "Algo de errado não deu certo"})
		} else {
			fmt.Println("Results All: ", results)
		}
		ctx.JSON(context.Map{"response": results})
		
	})

	app.Handle("GET", "/user/{msisdn: string}", func(ctx context.Context) {
		msisdn := ctx.Params().Get("msisdn")
		fmt.Println(msisdn)
	
		if(msisdn == "") {
			ctx.JSON(context.Map{"response": "Requisição incorreta, você deve passar o msisdn do usuário!"})
		} else {
			result := &User{}
			err = c.Find(bson.M{"msisdn": msisdn}).One(&result)
			if err != nil {
				ctx.JSON(context.Map{"response": err.Error()})
			} else {
				ctx.JSON(context.Map{"response": result});
			}
		}
	})

	app.Handle("POST", "/user", func(ctx context.Context) {
		params := &User{}
		err := ctx.ReadJSON(params)
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
		} else {
			params.LastUpdate = time.Now()
			err := c.Insert(params)
			if err != nil {
				ctx.JSON(context.Map{"response": err.Error()})
			} else {
				fmt.Println("Successfully inserted into database")
				result := User{}
				err = c.Find(bson.M{"msisdn": params.Msisdn}).One(&result)
				if err != nil {
					ctx.JSON(context.Map{"response": err.Error()})
				}
				ctx.JSON(context.Map{"response": "Usuário criado com sucesso!", "message": result})
			}
		}

	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}