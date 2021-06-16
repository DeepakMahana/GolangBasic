package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	deepak := person{
		firstName: "Deepak",
		lastName:  "Mahana",
		contact: contactInfo{
			email:   "deepak.mahana@gmail.com",
			zipCode: 80800,
		},
	}

	// Get the reference of the person
	// deepakPointer := &deepak
	// deepakPointer.updateName("Dee")

	// shortcut
	deepak.updateName("Dee")

	deepak.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
