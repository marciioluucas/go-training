package repository

import (
	"gopkg.in/mgo.v2/bson"
	"novaxs.com/go-training/iris-hello-world/src/models"
	"novaxs.com/go-training/iris-hello-world/src/utils"
	"novaxs.com/go-training/iris-hello-world/src/utils/database"
)

type UserRepository struct {
	AbstractRepository
}

func (userRepository *UserRepository) New() {
	db, _, err := database.Connect()
	userRepository.database = db
	if err != nil {
		panic(utils.Exception{}.New("Erro ao conectar ao banco de dados, `${err.Error()}`"))
	}

	userRepository.table = db.C("users")
}

func (userRepository *UserRepository) find(login string) *models.User {
	if login == "" {
		panic(utils.Exception{}.New("O login não pode ser vazio na busca de usuários por logins."))
	} else {
		result := &models.User{}
		err := userRepository.table.Find(bson.M{"login": login}).One(&result)
		if err != nil {
			panic(utils.Exception{}.New(err.Error()))
		} else {
			return result
		}
	}
}
