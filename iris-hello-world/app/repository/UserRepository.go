package repository

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"novaxs.com/go-training/iris-hello-world/app/models"
	"novaxs.com/go-training/iris-hello-world/app/utils"
	"novaxs.com/go-training/iris-hello-world/app/utils/database"
	"time"
)

type UserRepository struct {
	IRepository
	AbstractRepository
}

func (userRepository *UserRepository) Create(object models.User) *models.User {
	params := &object
	params.LastUpdate = time.Now()
	byteArray, err := params.GeneratePassword(params.Password)
	//n := bytes.IndexByte(byteArray, 0)
	params.Password = string(byteArray)
	e := userRepository.table.Insert(params)
	if e != nil {
		panic(e.Error())
	}

	fmt.Println("Usuario inserido com sucesso!")
	result := &models.User{}
	err = userRepository.table.Find(bson.M{"login": params.Login}).One(&result)
	if err != nil {
		panic(err.Error())
	}
	return result
}

func (userRepository *UserRepository) Retrieve() []models.User {
	results := []models.User{}
	err := userRepository.table.Find(nil).All(&results)
	if err != nil {
		panic(utils.Exception{}.New("Erro ao tentar buscar usuarios"))
	} else {
		fmt.Println("Results All: ", results)
	}
	return results
}

func (userRepository *UserRepository) Update(object interface{}) {
	panic("implement me")
}

func (userRepository *UserRepository) Delete(uniqueParam string) {
	panic("implement me")
}

func (userRepository *UserRepository) New() *UserRepository {
	db, session, err := database.Connect()
	if err != nil {
		panic(utils.Exception{}.New("Erro ao conectar ao banco de dados, `${err.Error()}`"))
	}
	index := mgo.Index{
		Key:        []string{"login"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	userRepository.table = db.C("users")
	err = userRepository.table.EnsureIndex(index)
	userRepository.database = db
	userRepository.session = &session
	return userRepository
}

func (userRepository *UserRepository) Find(login string) *models.User {
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
