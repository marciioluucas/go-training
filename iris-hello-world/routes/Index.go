package routes

import "github.com/kataras/iris"

func Index(app *iris.Application) {
	UserRoutes(app)
}
