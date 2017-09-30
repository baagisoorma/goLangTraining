package main

import "fmt"

func main() {
	switch "Savi" {
	case "Savi":
		fmt.Println("Kidda Savi!")
	case "Nav":
		fmt.Println("Kidda Nav!")
	case "Preet":
		fmt.Println("Kidda Preet!")
	default:
		fmt.Println("Hello no one:(")
	}
}