package controllers

import (
	"github.com/kataras/iris"
	"novaxs.com/go-training/iris-hello-world/src/utils"
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

}

func (uc UserController) put() {

}

func (uc UserController) delete() {

}
