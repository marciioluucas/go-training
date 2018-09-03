package repository

import "novaxs.com/go-training/iris-hello-world/app/models"

type IRepository interface {
	Create(object models.IModel)
	Retrieve()
	Update(object interface{})
	Delete(uniqueParam string)
	Find(uniqueParam string)
}
