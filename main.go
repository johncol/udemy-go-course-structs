package main

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
