package main

import "fmt"

type cat struct{}

func (c *cat) say() {
	fmt.Println("meow")
}

func main() {
	cat := &cat{}
	cat.say()
}
