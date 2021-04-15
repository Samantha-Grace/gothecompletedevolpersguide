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
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 560899,
		},
	}
	fmt.Printf("%+v\n", jim)
}

//embedded struct inside another
