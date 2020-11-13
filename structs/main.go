package main

import (
	"fmt"
)

type contactInformation struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInformation
}

func main() {
	var steve person
	steve.firstName = "Steve"
	steve.lastName  = "Jobs"
	fmt.Printf("The person is %+v", steve)
	
	jim := person {
		firstName: "Jim",
		lastName : "Carrey",
		contact  : contactInformation {
			email  : "jim@carrey.com",
			zipCode: 57,
		},
	}
	
	jim.updateName("Edwin")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}