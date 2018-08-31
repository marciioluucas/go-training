package models

import "fmt"

type Person struct {
    Name string;
    Email string;
}
func (p Person) PrintPerson() {
	fmt.Println(p.Name + ", " + p.Email);
}