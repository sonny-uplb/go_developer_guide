package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
