package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func (p person) getEmail() string {
	return p.email
}

func (p person) print() {
	fmt.Println()
	fmt.Printf("%+v", p)
}

// address as receiver
func (p *person) updateName(new string) {
	// update the value of the address
	(*p).firstName = new
}

func main() {
	var nof person
	// +v will add the fields' names
	fmt.Printf("%+v", nof)
	fmt.Println(nof)

	// define
	avishag := person{firstName: "Avishag", lastName: "Sahar", age: 25}
	fmt.Println(avishag)

	// update
	avishag.firstName = "Vishi"
	fmt.Println(avishag)

	jim := person{
		firstName: "Jim",
		lastName:  "Mor",
		age:       30,
		contactInfo: contactInfo{
			email:   "jimjim@jimjim.com",
			zipCode: 101010,
		},
	}
	fmt.Printf("%+v", jim.getEmail())
	jim.print()

	// update by pointer
	// & will give us the address
	jimPointer := &jim
	// update the address prop
	jimPointer.updateName("Jimmy")
	jim.print()

	// shortcuts for updates w/ pointer
	// auto compile
	jim.updateName("JimmyHoo")
	jim.print()

	// slice is actually contains a pointer
	// refernce types:
	// slice

	// value types:
	// native
	// structs

}
