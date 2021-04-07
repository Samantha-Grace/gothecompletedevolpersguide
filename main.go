package main

import (
	"fmt"
)

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
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 560899,
		},
	}
	fmt.Printf("%+v", jim)
}

//embedded struct inside another
