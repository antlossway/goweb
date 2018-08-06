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

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Toto"

	fmt.Println(alex)         //{Alex Toto}
	fmt.Printf("%+v\n", alex) //print all fields of struct: {firstName:Alex lastName:Toto}

	jim := person{
		firstName: "Jim",
		lastName:  "Job",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345, //multiline value of struct need to put comma on each line
		}, //multiline value of struct need to put comma on each line
	}

	jimPointer := &jim //jim is a reference to the actual value of struct
	jimPointer.updateName("jimmy")
	jim.print()

	fmt.Println("shortcut version")
	jim.updateName("jimmy2")
	jim.print()

}
