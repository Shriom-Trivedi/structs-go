package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	var name person
	name.firstName = "John"
	name.lastName = "Doe"
	name.contact.email = "John@doe.com"
	name.contact.zipCode = 233434

	fmt.Println(name)
}
