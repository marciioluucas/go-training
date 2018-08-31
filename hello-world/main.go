package main

import (
	"novaxs.com/hello-world/models"
)

func main() {
	person := models.Person{};
	person.Name = "Marcio Lucas";
	person.Email = "marciioluucas@gmail.com"
	person.PrintPerson()
}
