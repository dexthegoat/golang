package main

import "fmt"

// value type
// struct被传入一个func的时候 会copy一个整个的struct
// modify的时候不改变原来的
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo 像ES6 Object key val一样
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	var alex2 person
	fmt.Printf("%+v", alex2)
	alex2.firstName = "Alex2"
	alex2.lastName = "Anderson2"
	fmt.Printf("%+v", alex2)

	alex3 := person{
		firstName: "Alex3",
		lastName:  "Anderson3",
		contact: contactInfo{
			email:   "sadasd@gmail.com",
			zipCode: 07310,
		},
	}
	alex3.print()
	// &var Give me the memory address of var sitting at
	// alex3Pointer := &alex3 不写这行的话 go自动帮我们处理
	alex3.updateName("Jimmy")
	alex3.print()
}

// *person means we're working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// p是alex3的copy, so we're modifying a alex3 copy, can't do in this way
	// p.firstName = newFirstName
	// *pointerToPerson means we wanna manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// turn address into value with *address
// turn value into address with &value
