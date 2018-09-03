package controllers

import (
	"github.com/kataras/iris"
	"novaxs.com/go-training/iris-hello-world/app/models"
	"novaxs.com/go-training/iris-hello-world/app/repository"
	"novaxs.com/go-training/iris-hello-world/app/utils"
)

type UserController struct {
	app *iris.Application
	err error
}

func (uc UserController) setApp(app *iris.Application) {
	uc.app = app
}

func (uc UserController) get() {
	if uc.app == nil {
		e := utils.Exception{}.New("Aplicação não definida.")
		panic(e.GetMessage())
	}

}

func (uc UserController) post() {
	repo := repository.UserRepository{}.New()
	user := models.User{}
	user.Password = "dasdasd"
	repo.Create(user)
}

func (uc UserController) put() {

}

func (uc UserController) delete() {

}
