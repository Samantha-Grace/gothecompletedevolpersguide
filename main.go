package main

import "fmt"

func main() {
	colors := make(map[int]string)

	colors[10] = "fff8767"

	delete(colors, 10)

	fmt.Println(colors)
}

/*
func main() {
	colors := make(map[int]string)

	colors[10] = "fff8767"

	fmt.Println(colors)
}

func main() {
	colors := make(map[string]string)

	colors["white"] = "fff8767"

	fmt.Println(colors)
}

func main() {
	var colors map[string]string
	fmt.Println(colors)
}

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "4b6f45",
	}
	fmt.Println(colors)
}
*/
