package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}

func main() {
	person := person{
		firstName: "Michael",
		lastName: "Jackson",
		contactInfo: contactInfo{
			email: "michael@mail.com",
		},
	}

	person.print()

	person.updateName("Lorens")
	person.print()
}
