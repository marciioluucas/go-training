package controllers

type IController interface {
	get()
	post()
	put()
	delete()
}
