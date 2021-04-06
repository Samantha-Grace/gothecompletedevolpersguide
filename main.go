package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	//updating structs
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex)
}

/*
type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)
}


//second way

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
*/
